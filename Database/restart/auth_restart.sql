-- CLEAR TABLES
delete from accounts;
delete from role_permission;
delete from roles;
delete from permissions;

-- IMPORT DATA

INSERT INTO public.permissions (id,"name") VALUES
	 ('1de59064-e6ee-497c-8bb3-872d3f3e3607','Judge_crud'),
	 ('fe61d46d-ff68-4597-a794-1cc1290c7df6','SportsOrg_rud'),
	 ('0578facc-ced7-413a-bc9a-51222be4c178','Contestant_crud'),
	 ('22696b84-f022-4df1-beb6-976269a27bbe','Competition_cud'),
	 ('f5c0f18b-a724-4919-9eeb-8c5b7981a2c6','Application_crud'),
	 ('68bf18df-b44f-4ad8-806b-564c017c7ff9','Schedule_crud'),
	 ('2ef0b97e-4b72-49d0-842c-28538e75f88d','JudgingPanel_crud'),
	 ('421fbadf-95c6-481a-8f8b-63f7380268fa','LiveJudge_r'),
	 ('d854a572-4d14-47a5-80dc-499a5da0579a','LiveContestant_r'),
	 ('729ec6d9-47f6-4a00-950e-aabf5f9b41a1','Score_c'),
	 ('7b76338b-1db0-4ca1-b24f-860fa8520bdf','Score_r'),
	 ('dab42928-ef2a-47c5-9dd1-57522096121b','LiveSchedule_cru'),
	 ('feb4fba6-200a-46a9-a103-5a3a67cc04c0','WebSocket'),
	 ('d7424ccd-b2d5-4fb1-a8e7-5f11f1d68e71','ScoreBoard_r'),
	 ('5d889609-af22-4a87-8f45-4f948202f5e3','ScoreBoard_c');

	
INSERT INTO public.roles (id,"name") VALUES
	 ('ca8ec27c-aed4-4fc5-8d71-b3c4b34d917d','ADMIN'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','SPORTS_ORG'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','D_JUDGE'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','E_JUDGE');

INSERT INTO public.role_permission (role_id,permission_id) VALUES
	 ('26ab3170-938f-4468-8246-ab3602e4f016','1de59064-e6ee-497c-8bb3-872d3f3e3607'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','fe61d46d-ff68-4597-a794-1cc1290c7df6'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','0578facc-ced7-413a-bc9a-51222be4c178'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','22696b84-f022-4df1-beb6-976269a27bbe'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','f5c0f18b-a724-4919-9eeb-8c5b7981a2c6'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','68bf18df-b44f-4ad8-806b-564c017c7ff9'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','dab42928-ef2a-47c5-9dd1-57522096121b'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','2ef0b97e-4b72-49d0-842c-28538e75f88d'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','421fbadf-95c6-481a-8f8b-63f7380268fa'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','421fbadf-95c6-481a-8f8b-63f7380268fa'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','d854a572-4d14-47a5-80dc-499a5da0579a'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','d854a572-4d14-47a5-80dc-499a5da0579a'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','7b76338b-1db0-4ca1-b24f-860fa8520bdf'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','729ec6d9-47f6-4a00-950e-aabf5f9b41a1'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','729ec6d9-47f6-4a00-950e-aabf5f9b41a1'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','7b76338b-1db0-4ca1-b24f-860fa8520bdf'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','7b76338b-1db0-4ca1-b24f-860fa8520bdf'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','5d889609-af22-4a87-8f45-4f948202f5e3'),
	 ('7aec8ad4-ef7f-44b8-9645-82b16b7bed67','feb4fba6-200a-46a9-a103-5a3a67cc04c0'),
	 ('5a15ff04-b4ac-4f51-ba1b-e1efc74c3724','feb4fba6-200a-46a9-a103-5a3a67cc04c0'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','feb4fba6-200a-46a9-a103-5a3a67cc04c0'),
	 ('26ab3170-938f-4468-8246-ab3602e4f016','d7424ccd-b2d5-4fb1-a8e7-5f11f1d68e71');

INSERT INTO public.accounts (id,email,"password",role_id) VALUES
	 ('7ffc833e-322a-11ee-918f-0242ac120009','sdv@gmail.com','$2a$10$lT3fPKwtcNq33mE5A5wbE.BB1.dQrD4HyZiiXpb1dfHoXy/DHDcMe','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('c852b5e6-322a-11ee-918f-0242ac120009','mse@gmail.com','$2a$10$qY41vASL8HUYf7S6bPh3ZeV/a8lryhvO8zXzKOSs.moZ/FZxi2Qdu','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('f704dc99-322a-11ee-918f-0242ac120009','sports_rus@gmail.com','$2a$10$pizNZra/d9fiZoXcioXJNObr/Rb4OTSKYAhyIegEZ0S.qX6E.yyo.','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('58ada796-322b-11ee-918f-0242ac120009','agym@gmail.com','$2a$10$mLGZyN93mRf303a6bHWYceX3uRBX7PAG6QaZKM7xjCk.Agj9GVT9W','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('bab90d4d-33a0-11ee-918f-0242ac120009','bg1@gmail.com','$2a$10$pS7xY9ODjkkXLTCuUniDi.Z9Z..pqcyy5tUc1KS.0sSz2jtTEUZkq','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('eea5c376-33a0-11ee-918f-0242ac120009','bg2@gmail.com','$2a$10$b0nHD8Nd1B/3dpXrpIISjuauW.0HK4GtbiHc6mCuB3KCEKL9bNubK','26ab3170-938f-4468-8246-ab3602e4f016'),
	 ('97a9d610-33a1-11ee-918f-0242ac120009','hun2@gmail.com','$2a$10$mSvUhIH44lkF54w1Mj5FouO2pEtxFA97snWKvw/cyfCV64ysTCjiC','26ab3170-938f-4468-8246-ab3602e4f016');


