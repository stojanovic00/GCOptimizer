--CLEAR ALL TABLES


DELETE FROM apparatus_announcements;
DELETE FROM contestant_applications;
DELETE FROM age_categories;
DELETE FROM delegation_member_propositions;
DELETE FROM judge_applications;
DELETE FROM competitions;
DELETE FROM contestants; DELETE FROM judges;
DELETE FROM sports_organizations;
DELETE FROM addresses;
DELETE FROM delegation_member_applications;
--DELETE FROM delegation_member_positions ; we keep it because it will never change and i don't want to initialize it everytime
DELETE FROM delegation_members;
DELETE FROM team_compositions;

-- IMPORT DATA

INSERT INTO public.addresses (id,country,city,street,street_number) VALUES
	 ('4ea08568-3229-11ee-812d-0242ac12000a','Serbia','Novi Sad','Ignjata Pavlasa','2-4'),
	 ('80009b85-322a-11ee-812d-0242ac12000a','Serbia','Novi Sad','Ignjata Pavlasa','2-4'),
	 ('c862da31-322a-11ee-812d-0242ac12000a','Hungary','Budapest','Széchenyi István tér','9'),
	 ('f70893b2-322a-11ee-812d-0242ac12000a','Russia','Moscow','Leningradsky Prospekt ','39'),
	 ('58bf917d-322b-11ee-812d-0242ac12000a','United States','Phoenix','5th Avenue','78');


INSERT INTO public.sports_organizations (id,"name",email,phone_number,contact_person_full_name,competition_organising_privilege,address_id) VALUES
	 ('80009b81-322a-11ee-812d-0242ac12000a','Sokolsko Društvo Vojvodina','sdv@gmail.com','+3816540978','Geza Mikes',false,'80009b85-322a-11ee-812d-0242ac12000a'),
	 ('c862da2c-322a-11ee-812d-0242ac12000a','Magyar Sport Egyesület','mse@gmail.com','+36123456789','János Kovács',false,'c862da31-322a-11ee-812d-0242ac12000a'),
	 ('f70893b0-322a-11ee-812d-0242ac12000a','Спортивное Общество России','sports_rus@gmail.com','+79101234567','Ivan Petrov',false,'f70893b2-322a-11ee-812d-0242ac12000a'),
	 ('58bf9179-322b-11ee-812d-0242ac12000a','Arizona Gymnastics','agym@gmail.com','+14112378412','John Smith',false,'58bf917d-322b-11ee-812d-0242ac12000a');

INSERT INTO public.contestants (id,full_name,email,gender,position_id,image,sports_organization_id,date_of_birth) VALUES
	 ('00112248-3230-11ee-812d-0242ac12000a','Mina Đukić','mina.djukic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1998-06-12 01:00:00+02'),
	 ('0011224a-3230-11ee-812d-0242ac12000a','Jovana Petrović','jovana.petrovic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1997-09-05 01:00:00+02'),
	 ('0011224d-3230-11ee-812d-0242ac12000a','Mila Stanković','mila.stankovic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1996-08-20 01:00:00+02'),
	 ('0011224f-3230-11ee-812d-0242ac12000a','Jelena Nikolić','jelena.nikolic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','2000-05-15 01:00:00+02'),
	 ('00112250-3230-11ee-812d-0242ac12000a','Ivana Marković','ivana.markovic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1994-12-30 00:00:00+01'),
	 ('00112251-3230-11ee-812d-0242ac12000a','Maja Stojanović','maja.stojanovic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1999-10-25 01:00:00+02'),
	 ('00112252-3230-11ee-812d-0242ac12000a','Ana Janković','ana.jankovic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1995-11-02 00:00:00+01'),
	 ('00112253-3230-11ee-812d-0242ac12000a','Elena Đorđević','elena.djordjevic@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'80009b81-322a-11ee-812d-0242ac12000a','1998-03-18 00:00:00+01'),
	 ('00112254-3230-11ee-812d-0242ac12000a','Zsófia Nagy','zsofia.nagy@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1998-06-12 01:00:00+02'),
	 ('00112255-3230-11ee-812d-0242ac12000a','Lili Kovács','lili.kovacs@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1997-09-05 01:00:00+02'),
	 ('00112256-3230-11ee-812d-0242ac12000a','Emília Varga','emilia.varga@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1996-08-20 01:00:00+02'),
	 ('00112257-3230-11ee-812d-0242ac12000a','Zsuzsanna Farkas','zsuzsanna.farkas@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','2000-05-15 01:00:00+02'),
	 ('00112258-3230-11ee-812d-0242ac12000a','Borbála Balogh','borbala.balogh@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1994-12-30 00:00:00+01'),
	 ('00112259-3230-11ee-812d-0242ac12000a','Dóra Molnár','dora.molnar@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1999-10-25 01:00:00+02'),
	 ('0011225a-3230-11ee-812d-0242ac12000a','Katalin Papp','katalin.papp@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1995-11-02 00:00:00+01'),
	 ('0011225b-3230-11ee-812d-0242ac12000a','Réka Oláh','reka.olah@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'c862da2c-322a-11ee-812d-0242ac12000a','1998-03-18 00:00:00+01'),
	 ('0011225c-3230-11ee-812d-0242ac12000a','Анастасия Иванова','anastasia.ivanova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1998-06-12 01:00:00+02'),
	 ('0011225d-3230-11ee-812d-0242ac12000a','Екатерина Смирнова','ekaterina.smirnova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1997-09-05 01:00:00+02'),
	 ('0011225e-3230-11ee-812d-0242ac12000a','Ольга Федорова','olga.fedorova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1996-08-20 01:00:00+02'),
	 ('0011225f-3230-11ee-812d-0242ac12000a','Светлана Петрова','svetlana.petrova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','2000-05-15 01:00:00+02'),
	 ('00112260-3230-11ee-812d-0242ac12000a','Ирина Соколова','irina.sokolova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1994-12-30 00:00:00+01'),
	 ('00112261-3230-11ee-812d-0242ac12000a','Александра Кузнецова','alexandra.kuznetsova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1999-10-25 01:00:00+02'),
	 ('00112262-3230-11ee-812d-0242ac12000a','Елена Михайлова','elena.mikhailova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1995-11-02 00:00:00+01'),
	 ('00112263-3230-11ee-812d-0242ac12000a','Мария Волкова','maria.volkova@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'f70893b0-322a-11ee-812d-0242ac12000a','1998-03-18 00:00:00+01'),
	 ('00112264-3230-11ee-812d-0242ac12000a','Emma Johnson','emma.johnson@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1998-06-12 01:00:00+02'),
	 ('00112265-3230-11ee-812d-0242ac12000a','Olivia Williams','olivia.williams@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1997-09-05 01:00:00+02'),
	 ('00112266-3230-11ee-812d-0242ac12000a','Ava Smith','ava.smith@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1996-08-20 01:00:00+02'),
	 ('00112267-3230-11ee-812d-0242ac12000a','Sophia Brown','sophia.brown@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','2000-05-15 01:00:00+02'),
	 ('00112268-3230-11ee-812d-0242ac12000a','Isabella Miller','isabella.miller@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1994-12-30 00:00:00+01'),
	 ('00112269-3230-11ee-812d-0242ac12000a','Mia Davis','mia.davis@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1999-10-25 01:00:00+02'),
	 ('0011226a-3230-11ee-812d-0242ac12000a','Amelia Wilson','amelia.wilson@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1995-11-02 00:00:00+01'),
	 ('0011226b-3230-11ee-812d-0242ac12000a','Harper Anderson','harper.anderson@example.com',1,'8a7d23af-28a1-44b6-bfaf-5024f41821af',NULL,'58bf9179-322b-11ee-812d-0242ac12000a','1998-03-18 00:00:00+01');

INSERT INTO public.judges (id,full_name,email,gender,position_id,image,sports_organization_id,licence_type,licence_name) VALUES
	 ('f01a8257-32f1-11ee-812d-0242ac12000a','Jelena Petrović','jelena.petrovic@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'80009b81-322a-11ee-812d-0242ac12000a',1,'Professional Licence'),
	 ('f01a8258-32f1-11ee-812d-0242ac12000a','Milica Đukić','milica.djukic@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'80009b81-322a-11ee-812d-0242ac12000a',0,'Amateur Licence'),
	 ('f01a8259-32f1-11ee-812d-0242ac12000a','Ana Stanković','ana.stankovic@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'80009b81-322a-11ee-812d-0242ac12000a',1,'Level 1 Licence'),
	 ('f01a825a-32f1-11ee-812d-0242ac12000a','Jovana Nikolić','jovana.nikolic@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'80009b81-322a-11ee-812d-0242ac12000a',0,'Novice Licence'),
	 ('f01a825b-32f1-11ee-812d-0242ac12000a','Katalin Horváth','katalin.horvath@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'c862da2c-322a-11ee-812d-0242ac12000a',1,'Professional Licence'),
	 ('f01a825c-32f1-11ee-812d-0242ac12000a','Dóra Kovács','dora.kovacs@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'c862da2c-322a-11ee-812d-0242ac12000a',0,'Amateur Licence'),
	 ('f01a825d-32f1-11ee-812d-0242ac12000a','Eszter Szabó','eszter.szabo@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'c862da2c-322a-11ee-812d-0242ac12000a',1,'Level 1 Licence'),
	 ('f01a825e-32f1-11ee-812d-0242ac12000a','Zsuzsanna Varga','zsuzsanna.varga@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'c862da2c-322a-11ee-812d-0242ac12000a',0,'Novice Licence'),
	 ('f01a825f-32f1-11ee-812d-0242ac12000a','Anastasia Ivanova','anastasia.ivanova@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'f70893b0-322a-11ee-812d-0242ac12000a',1,'Professional Licence'),
	 ('f01a8260-32f1-11ee-812d-0242ac12000a','Ekaterina Petrova','ekaterina.petrova@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'f70893b0-322a-11ee-812d-0242ac12000a',0,'Amateur Licence'),
	 ('f01a8261-32f1-11ee-812d-0242ac12000a','Marina Sokolova','marina.sokolova@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'f70893b0-322a-11ee-812d-0242ac12000a',1,'Level 1 Licence'),
	 ('f01a8262-32f1-11ee-812d-0242ac12000a','Olga Popova','olga.popova@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'f70893b0-322a-11ee-812d-0242ac12000a',0,'Novice Licence'),
	 ('f01a8263-32f1-11ee-812d-0242ac12000a','Emily Johnson','emily.johnson@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'58bf9179-322b-11ee-812d-0242ac12000a',1,'Professional Licence'),
	 ('f01a8264-32f1-11ee-812d-0242ac12000a','Jessica Williams','jessica.williams@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'58bf9179-322b-11ee-812d-0242ac12000a',0,'Amateur Licence'),
	 ('f01a8265-32f1-11ee-812d-0242ac12000a','Sarah Brown','sarah.brown@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'58bf9179-322b-11ee-812d-0242ac12000a',1,'Level 1 Licence'),
	 ('f01a8266-32f1-11ee-812d-0242ac12000a','Elizabeth Miller','elizabeth.miller@gmail.com',1,'a7934c7f-4293-4877-90f1-6fc7f784185c',NULL,'58bf9179-322b-11ee-812d-0242ac12000a',0,'Novice Licence');
