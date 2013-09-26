-- Table: users
-- DROP TABLE users;
CREATE TABLE users
(
  uid serial NOT NULL,
  uname character varying(200),  --用户名
  passwd character(40), --用户密码
  utype smallint, -- 用户类型
  ulevel integer, -- 用户级别
  CONSTRAINT uid PRIMARY KEY (uid)
);
--ALTER TABLE users OWNER TO frm;
CREATE INDEX uname ON users USING btree (uname COLLATE pg_catalog."zh_CN.utf8");
INSERT INTO users (uname, passwd, utype, ulevel) VALUES ('root', '7c4a8d09ca3762af61e59520943dc26494f8941b', 1, 100);


-- Table: userType
CREATE TABLE userType
(
	utid serial NOT NULL,
	utname varchar(100),
	CONSTRAINT utid PRIMARY KEY (utid)
);
--ALTER TABLE userType OWNER TO frm;
CREATE INDEX utname ON userType USING btree (utname COLLATE pg_catalog."zh_CN.utf8");
INSERT INTO userType VALUES (1, '管理员');

-- Table: resources
-- DROP TABLE resources;
CREATE TABLE resources
(
  rid bigserial NOT NULL,
  rname character varying(1000),
  rtype smallint,
  rinfo text,
  rbtime bigint,
  rhashid character(40),
  CONSTRAINT rid PRIMARY KEY (rid)
);
--ALTER TABLE resources OWNER TO frm;
CREATE INDEX rhashid ON resources USING btree (rhashid COLLATE pg_catalog."default");


-- Table: resourceType
-- DROP TABLE resourceType
CREATE TABLE resourceType
(
	rtid serial NOT NULL,
	rtname varchar(100),
	CONSTRAINT rtid PRIMARY KEY (rtid)
);
--ALTER TABLE resources OWNER TO frm;
CREATE INDEX rtname ON resourceType USING btree (rtname COLLATE pg_catalog."zh_CN.utf8");
INSERT INTO resourceType VALUES (1, '图书');
