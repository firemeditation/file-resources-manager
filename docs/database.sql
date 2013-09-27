-- 机构表
create table units (
	unitid serial NOT NULL,
	unitname char(200),
	powerlevel json,  -- 使用json的语法存放权限等级值
	CONSTRAINT unitid PRIMARY KEY (unitid)
);
insert into units VALUES (1, '管理', '{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}');

-- 组表
create table groups (
	groupid serial not null,
	groupname char(200),
	powerlevel json,
	constraint groupid primary key (groupid)
);
insert into groups values (1, '管理', '{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}');

-- Table: users
CREATE TABLE users
(
  uid serial NOT NULL,
  uname char(200),  --用户名
  passwd char(40), --用户密码
  unitid int not null default 0, -- 用户所属机构
  groupid int not null default 0, -- 用户所属组
  powerlevel json,
  CONSTRAINT uid PRIMARY KEY (uid)
);
CREATE INDEX uname ON users USING btree (uname COLLATE pg_catalog."zh_CN.utf8");
ALTER TABLE users ADD FOREIGN KEY (unitid) REFERENCES units (unitid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE users ADD FOREIGN KEY (groupid) REFERENCES groups (groupid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
INSERT INTO users (uname, passwd, unitid, groupid, powerlevel) VALUES ('root', '7c4a8d09ca3762af61e59520943dc26494f8941b', 1, 1,'{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}');


-- Table: resourceType
CREATE TABLE resourceType
(
	rtid serial NOT NULL,
	rtname char(100),
	powerlevel int default 1,
	CONSTRAINT rtid PRIMARY KEY (rtid)
);
CREATE INDEX rtname ON resourceType USING btree (rtname COLLATE pg_catalog."zh_CN.utf8");
INSERT INTO resourceType VALUES (1, '图书', 1);


-- Table: 资源聚集
CREATE TABLE resourceGroup
(
  rid bigserial NOT NULL,
  rname char(1000),
  rtid smallint NOT NULL default 0, -- 资源类型
  rinfo text,
  rbtime bigint,
  rhashid char(40),
  unitid int not null default 0,  -- 对应的机构
  powerlevel int default 1,
  uid int NOT NULL default 0, -- 最后操作用户
  CONSTRAINT rid PRIMARY KEY (rid)
);
CREATE INDEX rhashid ON resourceGroup USING btree (rhashid);
CREATE INDEX rg_unitid ON resourceGroup USING btree (unitid);
CREATE INDEX rname ON resourceGroup USING btree (rname COLLATE pg_catalog."zh_CN.utf8");
ALTER TABLE resourceGroup ADD FOREIGN KEY (uid) REFERENCES users (uid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (rtid) REFERENCES resourceType (rtid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (unitid) REFERENCES units (unitid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;



-- Table: resourceFile
CREATE TABLE resourceFile
(
	rfhashid char(40) NOT NULL,  -- 哈希值(通过时间、文件名、路径名、资源id等混合得出)
	rid bigint NOT NULL,
	rfname char(1000) NOT NULL,  -- 文件名
	rfextname char(50) NOT NULL DEFAULT '', -- 文件扩展名
	rfpath char(2000) NOT NULL, -- 文件存放位置相对路径完整名字
	rfsite int NOT NULL,  -- 文件位置，主要是在需要多块硬盘的地方，由服务器配置文件制定序号
	rfsize bigint NOT NULL DEFAULT 0,  -- 文件字节数
	rflasttime bigint,  -- 最后更新日期
	rfver int,  -- 版本，每改一次加一
	unitid int not null default 0,  -- 对应的机构
	powerlevel int default 1,
	uid int NOT NULL default 0, -- 最后操作用户
	CONSTRAINT rfhashid PRIMARY KEY (rfhashid)
);
CREATE INDEX rfrid ON resourceFile USING btree (rid);
CREATE INDEX rfsite ON resourceFile USING btree (rfsite);
CREATE INDEX rfname ON resourceFile USING btree (rfname COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX rfextname ON resourceFile USING btree (rfextname);
CREATE INDEX rfunitid ON resourceFile USING btree (unitid);
ALTER TABLE resourcefile ADD FOREIGN KEY (rid) REFERENCES resources (rid) ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE resourcefile ADD FOREIGN KEY (uid) REFERENCES users (uid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourcefile ADD FOREIGN KEY (unitid) REFERENCES units (unitid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
