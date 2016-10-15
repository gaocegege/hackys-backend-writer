CREATE DATABASE writer;
\connect writer
CREATE TABLE tag
(
	id int,
	tag varchar(255)
);

INSERT INTO "tag" (id, tag) VALUES (1, 'anger');
INSERT INTO "tag" (id, tag) VALUES (2, 'contempt');
INSERT INTO "tag" (id, tag) VALUES (3, 'disgust');
INSERT INTO "tag" (id, tag) VALUES (4, 'fear');
INSERT INTO "tag" (id, tag) VALUES (5, 'happiness');
INSERT INTO "tag" (id, tag) VALUES (6, 'neutral');
INSERT INTO "tag" (id, tag) VALUES (7, 'sadness');
INSERT INTO "tag" (id, tag) VALUES (8, 'surprise');
