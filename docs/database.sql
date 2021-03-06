-- 机构表
drop table IF EXISTS units CASCADE;
create table units (
	id serial NOT NULL,
	name char(200),
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	powerlevel json default '{}',  -- 使用json的语法存放权限等级值
	info text, -- 机构信息
	CONSTRAINT units_id PRIMARY KEY (id)
);
ALTER TABLE units ADD UNIQUE (name);
insert into units VALUES (1, '空机构', 0, '{}', '空置机构');
insert into units VALUES (2, '总管', 0, '{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}', '系统的中心管理者');
insert into units VALUES (3, '机构一', 0, '{"user":{"user":0, "unit":0, "group":0},"resource":{"origin":1}}', '某一个机构');

-- 组表
drop table IF EXISTS groups CASCADE;
create table groups (
	id serial not null,
	name char(200),
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	powerlevel json default '{}',
	info text,  -- 组信息
	constraint groups_id primary key (id)
);
ALTER TABLE groups ADD UNIQUE (name);
insert into groups values (1, '空组', 0, '{}', '空组');
insert into groups values (2, '管理', 0, '{"user":{"user":5, "unit":5, "group":5},"resource":{"origin":5}}', '管理组');
insert into groups values (3, '机构管理员', 0, '{"user":{"user":4, "unit":4, "group":4},"resource":{"origin":4}}', '某一个机构的管理员');
insert into groups values (4, '普通使用者', 0, '{"resource":{"origin":2}}', '某一个机构的普通使用者');
insert into groups values (5, '普通查看者', 0, '{"resource":{"origin":1}}', '某一个机构的普通查看者');

-- Table: users
drop table IF EXISTS users CASCADE;
CREATE TABLE users
(
  id serial NOT NULL,
  name char(200),  --用户名
  passwd char(40), --用户密码
  units_id int not null default 1, -- 用户所属机构
  groups_id int not null default 1, -- 用户所属组
  expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
  powerlevel json default '{}',
  CONSTRAINT uid PRIMARY KEY (id)
);
ALTER TABLE users ADD UNIQUE (name);
CREATE INDEX name ON users USING btree (name COLLATE pg_catalog."zh_CN.utf8");
ALTER TABLE users ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE users ADD FOREIGN KEY (groups_id) REFERENCES groups (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
INSERT INTO users (id, name, passwd, units_id, groups_id) VALUES (1, 'nobody', '0000000000000000000000000000000000000000', 1, 1);
INSERT INTO users (id, name, passwd, units_id, groups_id) VALUES (2, 'root', '7c4a8d09ca3762af61e59520943dc26494f8941b', 2, 2);
INSERT INTO users (id, name, passwd, units_id, groups_id) VALUES (3, 'admin1', '7c4a8d09ca3762af61e59520943dc26494f8941b', 3, 3);
INSERT INTO users (id, name, passwd, units_id, groups_id) VALUES (4, 'admin2', '7c4a8d09ca3762af61e59520943dc26494f8941b', 3, 4);
INSERT INTO users (id, name, passwd, units_id, groups_id) VALUES (5, 'user1', '7c4a8d09ca3762af61e59520943dc26494f8941b', 3, 5);


-- Table: resourceType
drop table IF EXISTS resourceType CASCADE;
CREATE TABLE resourceType
(
	id serial NOT NULL,
	name char(100),
	powerlevel int default 1,
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	info text,
	CONSTRAINT rtid PRIMARY KEY (id)
);
CREATE INDEX rt_name ON resourceType USING btree (name COLLATE pg_catalog."zh_CN.utf8");
INSERT INTO resourceType VALUES (1, '无分类', 1, 0, '无分类');
INSERT INTO resourceType VALUES (2, '图书', 1, 0, '图书分类');
INSERT INTO resourceType VALUES (3, '杂志', 1, 0, '杂志分类');


-- Table: 资源聚集
drop table IF EXISTS resourceGroup CASCADE;
CREATE TABLE resourceGroup
(
  hashid char(40),
  name char(1000),
  rt_id smallint NOT NULL default 1, -- 资源类型
  info text,
  btime bigint, -- 创建时间
  derivative char(40) not null default '0000000000000000000000000000000000000000',  -- 衍生自哪个resourceGroup，这是衍生的hashid，如果没有衍生就是null
  units_id int not null default 1,  -- 对应的机构
  powerlevel int default 1,
  users_id int NOT NULL default 1, -- 最后操作用户
  expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
  metadata json not null default '{}',
  CONSTRAINT rgkey_hashid PRIMARY KEY (hashid)
);
CREATE INDEX rg_hashid ON resourceGroup USING btree (hashid);
CREATE INDEX rg_unitid ON resourceGroup USING btree (units_id);
CREATE INDEX rg_derivative ON resourceGroup USING btree (derivative);
CREATE INDEX rg_name ON resourceGroup USING btree (name COLLATE pg_catalog."zh_CN.utf8");

insert into resourceGroup (hashid, name, info, btime) values ('0000000000000000000000000000000000000000', '空资源聚集','空资源聚集',1);

ALTER TABLE resourceGroup ADD FOREIGN KEY (users_id) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (rt_id) REFERENCES resourceType (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (derivative) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD UNIQUE (hashid);


-- 资源聚集状态
drop table IF EXISTS resourceGroupStatus  CASCADE;
create table resourceGroupStatus
(
	hashid char(40) not null,
	status1 int not null default 0, -- 资源聚集的原始文件个数
	status2 int not null default 0,
	status3 int not null default 0,
	status4 int not null default 0,
	status5 int not null default 0,
	status6 int not null default 0,
	status7 int not null default 0,
	status8 int not null default 0,
	status9 int not null default 0,
	CONSTRAINT rgs_hashid PRIMARY KEY (hashid)
);
ALTER TABLE resourceGroupStatus ADD FOREIGN KEY (hashid) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;


-- Table: 资源条目
drop table IF EXISTS resourceItem CASCADE;
CREATE TABLE resourceItem
(
	hashid char(40) NOT NULL,  -- 哈希值(通过时间、文件名、路径名、资源id等混合得出)
	name char(1000) NOT NULL,  -- 文件名
	ritype int not null default 0,  -- 资源条目类型：1为file，2为text，之后增加就继续
	lasttime bigint,  -- 最后更新日期
	version int not null default 1,  -- 版本，每改一次加一
	rg_hashid char(40) not null default '0000000000000000000000000000000000000000', -- 资源条目的原始聚集ID
	derivative char(40) not null default '0000000000000000000000000000000000000000',  -- 衍生自哪个resourceItem，这是衍生的hashid，如果没有衍生就是null
	units_id int not null default 1,  -- 对应的机构
	powerlevel int default 1,
	users_id int NOT NULL default 1, -- 最后操作用户
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	metadata json not null default '{}',
	CONSTRAINT ri_hashid PRIMARY KEY (hashid)
);
CREATE INDEX riname ON resourceItem USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX ritype ON resourceItem USING btree (ritype);
CREATE INDEX riunitid ON resourceItem USING btree (units_id);
CREATE INDEX ri_rg_hashid ON resourceItem USING btree (rg_hashid);
CREATE INDEX ri_rg_derivative ON resourceItem USING btree (derivative);

insert into resourceItem (hashid, name, lasttime) values ('0000000000000000000000000000000000000000', '空资源条目', 1);

ALTER TABLE resourceItem ADD FOREIGN KEY (rg_hashid) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE SET NULL;
ALTER TABLE resourceItem ADD FOREIGN KEY (users_id) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceItem ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceItem ADD FOREIGN KEY (derivative) REFERENCES resourceItem (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE resourceItem ADD UNIQUE (hashid);
  
-- Table: 资源文件 从资源条目继承
drop table if exists resourceFile cascade;
create table resourceFile (
	fname char(2000) ,  -- 文件名
	extname char(50) NOT NULL DEFAULT '', -- 文件扩展名
	opath char(2000) NOT NULL, -- 文件的原始相对路径
	fpath char(2000) NOT NULL, -- 文件存放位置相对路径完整名字
	fsite char(2000) NOT NULL,  -- 文件位置，主要是在需要多块硬盘的地方
	fsize bigint NOT NULL DEFAULT 0,  -- 文件字节数
	CONSTRAINT rf_hashid PRIMARY KEY (hashid)
) INHERITS (resourceItem);
CREATE INDEX rfsite ON resourceFile USING btree (fsite);
CREATE INDEX rfextname ON resourceFile USING btree (extname);
CREATE INDEX rthashid ON resourceFile USING btree (hashid);
CREATE INDEX rfname ON resourceFile USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX rfunitid ON resourceFile USING btree (units_id);
CREATE INDEX rf_rg_id ON resourceFile USING btree (rg_hashid);

  -- 文件的资源条目状态
drop table IF EXISTS resourceFileStatus  CASCADE;
create table resourceFileStatus
(
	hashid char(40) not null,
	status1 int not null default 0,
	status2 int not null default 0,
	status3 int not null default 0,
	status4 int not null default 0,
	status5 int not null default 0,
	status6 int not null default 0,
	status7 int not null default 0,
	status8 int not null default 0,
	status9 int not null default 0,
	CONSTRAINT rfs_hashid PRIMARY KEY (hashid)
);
ALTER TABLE resourceFileStatus ADD FOREIGN KEY (hashid) REFERENCES resourceFile (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;
  


-- Table: 资源文本 从资源条目继承
drop table IF EXISTS resourceText cascade;
create table resourceText (
	c_type char(20),  -- 资源条目的类型，比如“menu”之类的
	conent text,
	CONSTRAINT rtf_hashid PRIMARY KEY (hashid)
) INHERITS (resourceItem);
CREATE INDEX rtfhashid ON resourceText USING btree (hashid);
CREATE INDEX rtctype ON resourceText USING btree (c_type);
CREATE INDEX rtfname ON resourceText USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX rtfunitid ON resourceText USING btree (units_id);
CREATE INDEX rft_rg_id ON resourceText USING btree (rg_hashid);

  -- 文本的资源条目状态
drop table IF EXISTS resourceTextStatus  CASCADE;
create table resourceTextStatus
(
	hashid char(40) not null,
	status1 int not null default 0,
	status2 int not null default 0,
	status3 int not null default 0,
	status4 int not null default 0,
	status5 int not null default 0,
	status6 int not null default 0,
	status7 int not null default 0,
	status8 int not null default 0,
	status9 int not null default 0,
	CONSTRAINT rts_hashid PRIMARY KEY (hashid)
);
ALTER TABLE resourceTextStatus  ADD FOREIGN KEY (hashid) REFERENCES resourceText (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;
  


-- 资源聚集关系
drop table IF EXISTS resourceRelation CASCADE;
create table resourceRelation (
	quote_side char(40),  -- 引用方
	be_quote char(40),  -- 被引用方
	rr_type int  --聚集类型，1文件聚集（引用方指resourceGroup），2嵌入聚集（引用方指resourceItem）
);
create index rr_rgid on resourceRelation using btree (quote_side);
create index rr_rfhashid on resourceRelation using btree (be_quote);
ALTER TABLE resourceRelation ADD FOREIGN KEY (quote_side) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE resourceRelation ADD FOREIGN KEY (quote_side) REFERENCES resourceItem (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;
ALTER TABLE resourceRelation ADD FOREIGN KEY (be_quote) REFERENCES resourceItem (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;


-- 异步缓存全文索引 Async Cache Full Text Index
drop table if exists acfti cascade;
create table acfti (
	key_word char(100), --关键词
	uid int, -- 所在机构的id
	hashid char(40), --是哪个HashId
	htype int  --对应为1：rg、2：rf、3：rt
);
CREATE INDEX acftikeywordindex ON acfti USING btree (key_word);
CREATE INDEX acftiuidindex ON acfti USING btree (uid);
CREATE INDEX acftihashidindex ON acfti USING btree (hashid);
CREATE INDEX acftihtypeindex ON acfti USING btree (htype);
ALTER TABLE acfti ADD FOREIGN KEY (uid) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
