DROP TABLE IF EXISTS users;
CREATE TABLE IF NOT EXISTS users(
    "id" VARCHAR(32) NOT NULL PRIMARY KEY,
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(255) NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP
);
INSERT INTO users ("id", "email", "password", "created_at", "updated_at", "deleted_at") VALUES
	 ('2BUrOvAvqHbwSIxe3iAHYJekMmP','mail1@mail.com','$2a$08$r1K9stWxsqiv/fTayPvMIuVq1ReqZJf.kqwaWQE1FZ.FKOquhKExS','2022-07-04 21:51:16.008221','2022-07-04 21:51:16.008221',NULL),
	 ('2BUrQKjp7IFM2aTUx70yCLtFPot','mail2@mail.com','$2a$08$OW8SWEjJkZ99Zr1Xs8KxZ.3am/9dBxogijsdePMRA.QS4jsWXoVga','2022-07-04 21:51:26.614677','2022-07-04 21:51:26.614677',NULL),
	 ('2BUrSEq4uEtmFmp7Y9x25lWtK5n','mail3@mail.com','$2a$08$ufUb7fK2bDUdyFjmCf7QBOQmc8gOgiO8K5dC3GqrGVE/CMu5MrpJ2','2022-07-04 21:51:42.276974','2022-07-04 21:51:42.276974',NULL),
	 ('2BUrTGveAFtqNANj13BRa9vpjkK','mail4@mail.com','$2a$08$sPRsYQd2vsBYPf8MzPENNO7fmFze918PmmWWv0E/B8qLhrMkmYjHu','2022-07-04 21:51:50.479185','2022-07-04 21:51:50.479185',NULL),
	 ('2BUrU5g2eMS77sVcxj5sMZYEe1u','mail5@mail.com','$2a$08$wx13LVym4.8gTit3t/uTM.k1I63TPvLePaEg.0lDfmuiNxcibuZAK','2022-07-04 21:51:56.251079','2022-07-04 21:51:56.251079',NULL),
	 ('2BUrUreTPDrnl11WXEiweIXLpVk','mail6@mail.com','$2a$08$Zo4CQyoZCg8OwiGSBTf3xuIdxdjk8zSJhQpM8YGxeLa7uHvgTWMG2','2022-07-04 21:52:03.678638','2022-07-04 21:52:03.678638',NULL),
	 ('2BUrVYSFjrqrkp3KYwJe0rf1jws','mail7@mail.com','$2a$08$OzBgM/MCUy2dyMv14qNgf.SwAXCBNClO3IP0aDvs4dIk..9qRDRYO','2022-07-04 21:52:08.29849','2022-07-04 21:52:08.29849',NULL),
	 ('2BUrWAcgnppkyjkBIo6CQofGnaD','mail8@mail.com','$2a$08$DSEOwWAJ5Hs4B036A4NPBeITXjlXW5AApM0pstdPIpkwUGQ5Wk9bK','2022-07-04 21:52:13.096474','2022-07-04 21:52:13.096474',NULL),
	 ('2BUrWvjrsu7HcGgBsVDDlbfzx6T','mail9@mail.com','$2a$08$qvm6oRSaJKggkBCAS6yz.ekRnpyw8S/PlTRfH84nCFxVIV32GO1si','2022-07-04 21:52:19.037826','2022-07-04 21:52:19.037826',NULL),
	 ('2BUrXUqkTwGwAiASan3XJY4paTS','mail10@mail.com','$2a$08$w52WWWb9K6NG0Nd..ijUW.lUBOoRKChggPf4bE9p9gNUMrJYT9t1G','2022-07-04 21:52:23.936662','2022-07-04 21:52:23.936662',NULL),
	 ('2BUrXwfsIPnV1W6L9vn67uc343v','mail11@mail.com','$2a$08$SD48MWTWC/tuFnM56y5hT..6EeIfTKCbtYPrU5reJEIbFgc5GeBZC','2022-07-04 21:52:27.66736','2022-07-04 21:52:27.66736',NULL),
	 ('2BUrYVYmKEjSZ8Jgg0j5Exknvo5','mail12@mail.com','$2a$08$zxBbG35Nu.s5siRRzSRxZOvzPxL1IQ.Qd8kzU.3XIh0Vys61krKvu','2022-07-04 21:52:32.098523','2022-07-04 21:52:32.098523',NULL),
	 ('2BUrZwSzZwlnJyKeEXjm3fK9uAw','mail13@mail.com','$2a$08$/fLpYxjqWs9QT17v4vjX9uHqMoZWmCoRRWSXPOp7bvX.O.fB2kqdS','2022-07-04 21:52:43.924148','2022-07-04 21:52:43.924148',NULL),
	 ('2BUramFRWWCHYqjsgUIHQqF0KWc','mail14@mail.com','$2a$08$h2Ye16QsLzp3dWkLR0I4Ou8Pp9ywh3D/og9yIMIguH0LZQddukmsy','2022-07-04 21:52:50.505252','2022-07-04 21:52:50.505252',NULL),
	 ('2BUrbLJH9zjoUXZxQmnlubm2B5d','mail15@mail.com','$2a$08$UQE7hAckLJcHpyJDBEW1p.P2kmb/2oL9HTAFaGzE1JvD0NqcoOr8O','2022-07-04 21:52:54.463897','2022-07-04 21:52:54.463897',NULL),
	 ('2BUrboVEknMyusATs4iavgDkUhV','mail16@mail.com','$2a$08$8k6oR3t3N9esZFE9doaV6ujThJOZAlM85lxwxpCoYuXp5JNU4FFrC','2022-07-04 21:52:58.075048','2022-07-04 21:52:58.075048',NULL),
	 ('2BUrcDnghkqpcjQuMadL2qHmDRI','mail17@mail.com','$2a$08$WKZD/ENBH0nv9K41lnARvOVv31v5mF5TVKb/JetGE7oazesyQAN/e','2022-07-04 21:53:01.423741','2022-07-04 21:53:01.423741',NULL),
	 ('2BUrcbF4Iekhn6SJYqR67XBKoZq','mail18@mail.com','$2a$08$PLEBgTi8OsapGHZ641i91ekCLv0uIJLnyPlTBneNC5BAiJlOdO9.u','2022-07-04 21:53:04.629812','2022-07-04 21:53:04.629812',NULL),
	 ('2BUrdGBTzh36O83TpQAwVsgyIcA','mail19@mail.com','$2a$08$yX4frEznsdgFeV4VqWqXmOdwYaAMfqscr88pYMyYJJoRvnBI1B.Sm','2022-07-04 21:53:09.554407','2022-07-04 21:53:09.554407',NULL),
	 ('2BUrdx5M7aUWZAoeSujnaslJlBF','mail20@mail.com','$2a$08$rvW5Wu2kk4mqrdIVGroEUeDdtBZXJ41xGpmAgCS1R8N0O.bI.FyDO','2022-07-04 21:53:15.278289','2022-07-04 21:53:15.278289',NULL)
;