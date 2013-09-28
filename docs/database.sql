-- 机构表
drop table IF EXISTS units CASCADE;
create table units (
	id serial NOT NULL,
	name char(200),
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	powerlevel json,  -- 使用json的语法存放权限等级值
	info text, -- 机构信息
	CONSTRAINT units_id PRIMARY KEY (id)
);
insert into units VALUES (1, '管理', 0, '{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}', '管理机构');

-- 组表
drop table IF EXISTS groups CASCADE;
create table groups (
	id serial not null,
	name char(200),
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	powerlevel json,
	info text,  -- 组信息
	constraint groups_id primary key (id)
);
insert into groups values (1, '管理', 0, '{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}', '管理组');

-- Table: users
drop table IF EXISTS users CASCADE;
CREATE TABLE users
(
  id serial NOT NULL,
  name char(200),  --用户名
  passwd char(40), --用户密码
  units_id int not null default 0, -- 用户所属机构
  groups_id int not null default 0, -- 用户所属组
  expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
  powerlevel json,
  CONSTRAINT uid PRIMARY KEY (id)
);
CREATE INDEX name ON users USING btree (name COLLATE pg_catalog."zh_CN.utf8");
ALTER TABLE users ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE users ADD FOREIGN KEY (groups_id) REFERENCES groups (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
INSERT INTO users (name, passwd, units_id, groups_id, powerlevel) VALUES ('root', '7c4a8d09ca3762af61e59520943dc26494f8941b', 1, 1,'{"user":{"user":1, "unit":1, "group":1},"resource":{"origin":1}}');


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
INSERT INTO resourceType VALUES (1, '图书', 1, 0, '图书分类');


-- Table: 资源聚集
drop table IF EXISTS resourceGroup CASCADE;
CREATE TABLE resourceGroup
(
  hashid char(40),
  name char(1000),
  rt_id smallint NOT NULL default 0, -- 资源类型
  info text,
  btime bigint, -- 创建时间
  derivative char(40),  -- 衍生自哪个resourceGroup，这是衍生的hashid，如果没有衍生就是null
  units_id int not null default 0,  -- 对应的机构
  powerlevel int default 1,
  users_id int NOT NULL default 0, -- 最后操作用户
  expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
  CONSTRAINT rgkey_hashid PRIMARY KEY (hashid)
);
CREATE INDEX rg_hashid ON resourceGroup USING btree (hashid);
CREATE INDEX rg_unitid ON resourceGroup USING btree (units_id);
CREATE INDEX rg_derivative ON resourceGroup USING btree (derivative);
CREATE INDEX rg_name ON resourceGroup USING btree (name COLLATE pg_catalog."zh_CN.utf8");
ALTER TABLE resourceGroup ADD FOREIGN KEY (users_id) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (rt_id) REFERENCES resourceType (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceGroup ADD FOREIGN KEY (derivative) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;



-- Table: 资源条目
drop table IF EXISTS resourceItem CASCADE;
CREATE TABLE resourceItem
(
	hashid char(40) NOT NULL,  -- 哈希值(通过时间、文件名、路径名、资源id等混合得出)
	name char(1000) NOT NULL,  -- 文件名
	lasttime bigint,  -- 最后更新日期
	verson int,  -- 版本，每改一次加一
	rg_hashid char(40), -- 资源条目的原始聚集ID
	derivative char(40),  -- 衍生自哪个resourceItem，这是衍生的hashid，如果没有衍生就是null
	units_id int not null default 0,  -- 对应的机构
	powerlevel int default 1,
	users_id int NOT NULL default 0, -- 最后操作用户
	expand int not null default 0, -- 扩展，默认没有扩展表，有扩展表则写明编号
	CONSTRAINT ri_hashid PRIMARY KEY (hashid)
);
CREATE INDEX riname ON resourceItem USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX riunitid ON resourceItem USING btree (units_id);
CREATE INDEX ri_rg_hashid ON resourceItem USING btree (rg_hashid);
CREATE INDEX ri_rg_derivative ON resourceItem USING btree (derivative);
ALTER TABLE resourceItem ADD FOREIGN KEY (rg_hashid) REFERENCES resourceGroup (hashid) ON UPDATE NO ACTION ON DELETE SET NULL;
ALTER TABLE resourceItem ADD FOREIGN KEY (users_id) REFERENCES users (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceItem ADD FOREIGN KEY (units_id) REFERENCES units (id) ON UPDATE NO ACTION ON DELETE SET DEFAULT;
ALTER TABLE resourceItem ADD FOREIGN KEY (derivative) REFERENCES resourceItem (hashid) ON UPDATE NO ACTION ON DELETE CASCADE;

-- Table: 资源文件 从资源条目继承
drop table if exists resourceFile cascade;
create table resourceFile (
	fname char(2000) ,  -- 文件名
	extname char(50) NOT NULL DEFAULT '', -- 文件扩展名
	opath char(2000) NOT NULL, -- 文件的原始相对路径
	fpath char(2000) NOT NULL, -- 文件存放位置相对路径完整名字
	fsite int NOT NULL,  -- 文件位置，主要是在需要多块硬盘的地方，由服务器配置文件制定序号
	fsize bigint NOT NULL DEFAULT 0,  -- 文件字节数
	metadata json, -- 元数据
	CONSTRAINT rf_hashid PRIMARY KEY (hashid)
) INHERITS (resourceItem);
CREATE INDEX rfsite ON resourceFile USING btree (fsite);
CREATE INDEX rfextname ON resourceFile USING btree (extname);
CREATE INDEX rthashid ON resourceFile USING btree (hashid);
CREATE INDEX rfname ON resourceFile USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX rfunitid ON resourceFile USING btree (units_id);
CREATE INDEX rf_rg_id ON resourceFile USING btree (rg_hashid);


-- Table: 资源文本 从资源条目继承
drop table IF EXISTS resourceText cascade;
create table resourceText (
	conent text,
	metadata json, -- 元数据
	CONSTRAINT rtf_hashid PRIMARY KEY (hashid)
) INHERITS (resourceItem);
CREATE INDEX rtfhashid ON resourceText USING btree (hashid);
CREATE INDEX rtfname ON resourceText USING btree (name COLLATE pg_catalog."zh_CN.utf8");
CREATE INDEX rtfunitid ON resourceText USING btree (units_id);
CREATE INDEX rft_rg_id ON resourceText USING btree (rg_hashid);


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
