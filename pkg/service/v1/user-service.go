package v1

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	v1 "github.com/aerodinamicat/thisisme01/pkg/api/v1"
	"github.com/golang/protobuf/ptypes"
	_ "github.com/lib/pq"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	API_VERSION = "v1"
	HASH_COST   = 8
)

type userServiceServer struct {
	db *sql.DB
	v1.UnimplementedUserServiceServer
}

func NewUserServiceServer(db *sql.DB) v1.UserServiceServer {
	return &userServiceServer{db: db}
}

func (s *userServiceServer) checkAPIVersion(api string) error {
	if len(api) > 0 {
		if API_VERSION != api {
			return status.Errorf(codes.Unimplemented, "Unsupported API version: expected '%s' but got '%s'", API_VERSION, api)
		}
	}
	return nil
}
func (s *userServiceServer) connect(ctx context.Context) (*sql.Conn, error) {
	c, err := s.db.Conn(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to connect to database: '%v'", err.Error())
	}
	return c, nil
}

func (s *userServiceServer) Create(ctx context.Context, req *v1.CreateRequest) (*v1.CreateResponse, error) {
	/*if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}*/

	c, err := s.connect(ctx)
	if err != nil {
		//* Error de conexión al servidor de base de datos
		return nil, err
	}
	defer c.Close()

	generatedId, err := ksuid.NewRandom()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate random id: '%v'", err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().GetPassword()), HASH_COST)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash given user password: '%v'", err.Error())
	}

	querySentence := "INSERT INTO users (id,email,password) VALUES ($1,$2,$3)"
	_, err = c.ExecContext(ctx, querySentence, generatedId.String(), req.GetUser().GetEmail(), string(hashedPassword))
	if err != nil {
		return nil, err
	}

	return &v1.CreateResponse{
		Id: generatedId.String(),
	}, nil
}
func (s *userServiceServer) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	/*if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}*/

	c, err := s.connect(ctx)
	if err != nil {
		//* Error de conexión al servidor de base de datos
		return nil, err
	}
	defer c.Close()

	querySentence := "SELECT id, email, password, created_at, updated_at, deleted_at FROM users WHERE id = $1"
	rows, err := c.QueryContext(ctx, querySentence, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to select from DB: '%v'", err.Error())
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, status.Errorf(codes.Unknown, "Failed to retreive data from DB: '%v'", err.Error())
		}
		return nil, status.Errorf(codes.NotFound, "User with 'id' = '%s' is not found in DB: '%v'", req.GetId(), err.Error())
	}

	var user v1.User
	var createdAt, updatedAt, deletedAt sql.NullTime
	if err := rows.Scan(&user.Id, &user.Email, &user.Password, &createdAt, &updatedAt, &deletedAt); err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to retrieve field values User row: '%v'", err.Error())
	}

	if createdAt.Valid {
		if user.CreatedAt, err = ptypes.TimestampProto(createdAt.Time); err != nil {
			return nil, status.Errorf(codes.Unknown, "'CreatedAt' field has invalid format: '%v'", err.Error())
		}
	} else {
		user.CreatedAt = nil
	}
	if updatedAt.Valid {
		if user.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time); err != nil {
			return nil, status.Errorf(codes.Unknown, "'UpdatedAt' field has invalid format: '%v'", err.Error())
		}
	} else {
		user.UpdatedAt = nil
	}
	if deletedAt.Valid {
		if user.DeletedAt, err = ptypes.TimestampProto(deletedAt.Time); err != nil {
			return nil, status.Errorf(codes.Unknown, "'DeletedAt' field has invalid format: '%v'", err.Error())
		}
	} else {
		user.DeletedAt = nil
	}

	if rows.Next() {
		return nil, status.Errorf(codes.Unknown, "Found multiple User rows with 'id' = '%s': '%v'", req.GetId(), err.Error())
	}

	return &v1.GetResponse{
		User: &user,
	}, nil
}
func (s *userServiceServer) Update(ctx context.Context, req *v1.UpdateRequest) (*v1.UpdateResponse, error) {
	/*if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}*/

	c, err := s.connect(ctx)
	if err != nil {
		//* Error de conexión al servidor de base de datos
		return nil, err
	}
	defer c.Close()

	user := req.GetUser()

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.GetUser().GetPassword()), HASH_COST)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash given user password: '%v'", err.Error())
	}
	user.Password = string(hashedNewPassword)
	updatedAt := time.Now()
	if user.UpdatedAt, err = ptypes.TimestampProto(updatedAt); err != nil {
		return nil, status.Errorf(codes.Unknown, "'UpdatedAt' field has invalid format: '%v'", err.Error())
	}

	querySentence := "UPDATE users SET email = $1, password = $2, updated_at = $3 WHERE id = $4"
	res, err := c.ExecContext(ctx, querySentence, user.Email, user.Password, updatedAt, user.Id)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to update user: '%v'", err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to retrieve rows affected: '%v'", err.Error())
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User with 'id' = '%s' is not found: '%v'", err.Error())
	}

	return &v1.UpdateResponse{
		RowsAffected: fmt.Sprintf("%d", rowsAffected),
	}, nil
}
func (s *userServiceServer) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.DeleteResponse, error) {
	/*if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}*/

	c, err := s.connect(ctx)
	if err != nil {
		//* Error de conexión al servidor de base de datos
		return nil, err
	}
	defer c.Close()

	querySentence := "DELETE FROM users WHERE id = $1"
	res, err := c.ExecContext(ctx, querySentence, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to delete user: '%v'", err.Error())
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to retrieve rows affected: '%v'", err.Error())
	}
	if rowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "User with 'id' = '%s' is not found: '%v'", err.Error())
	}

	return &v1.DeleteResponse{
		RowsAffected: fmt.Sprintf("%d", rowsAffected),
	}, nil
}
func (s *userServiceServer) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	/*if err := s.checkAPIVersion(req.Api); err != nil {
		return nil, err
	}*/

	c, err := s.connect(ctx)
	if err != nil {
		//* Error de conexión al servidor de base de datos
		return nil, err
	}
	defer c.Close()

	querySentence := "SELECT id, email, password, created_at, updated_at, deleted_at FROM users"
	rows, err := c.QueryContext(ctx, querySentence)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to select from DB: '%v'", err.Error())
	}
	defer rows.Close()

	var createdAt, updatedAt, deletedAt sql.NullTime
	list := []*v1.User{}
	for rows.Next() {
		user := &v1.User{}
		if err := rows.Scan(&user.Id, &user.Email, &user.Password, &createdAt, &updatedAt, &deletedAt); err != nil {
			return nil, status.Errorf(codes.Unknown, "Failed to retrieve field values User row: '%v'", err.Error())
		}

		if createdAt.Valid {
			if user.CreatedAt, err = ptypes.TimestampProto(createdAt.Time); err != nil {
				return nil, status.Errorf(codes.Unknown, "'CreatedAt' field has invalid format: '%v'", err.Error())
			}
		} else {
			user.CreatedAt = nil
		}
		if updatedAt.Valid {
			if user.UpdatedAt, err = ptypes.TimestampProto(updatedAt.Time); err != nil {
				return nil, status.Errorf(codes.Unknown, "'UpdatedAt' field has invalid format: '%v'", err.Error())
			}
		} else {
			user.UpdatedAt = nil
		}
		if deletedAt.Valid {
			if user.DeletedAt, err = ptypes.TimestampProto(deletedAt.Time); err != nil {
				return nil, status.Errorf(codes.Unknown, "'DeletedAt' field has invalid format: '%v'", err.Error())
			}
		} else {
			user.DeletedAt = nil
		}

		list = append(list, user)
	}

	if err := rows.Err(); err != nil {
		return nil, status.Errorf(codes.Unknown, "Failed to retreive data from DB: '%v'", err.Error())
	}

	return &v1.ListResponse{
		Users: list,
	}, nil
}
