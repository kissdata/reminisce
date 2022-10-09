-- 建表语句示例
-- 存储好友信息

-- 安装 uuid_generate_v4() 扩展函数
create extension "uuid-ossp"

--
-- 好友信息表
--
CREATE TABLE if not exists public.friend
(
  ts         timestamptz DEFAULT now(),
  friend_id  UUID        DEFAULT uuid_generate_v4(),
  name       VARCHAR(20) NOT NULL,
  first_know date        NULL,
  remark     VARCHAR     DEFAULT NULL,
  creator    VARCHAR(20) DEFAULT 'njupter',
  PRIMARY KEY (friend_id)
)with (oids = false);

COMMENT ON TABLE public.friend             IS '好友清单';
COMMENT ON COLUMN public.friend.ts         IS '创建时间';
COMMENT ON COLUMN public.friend.friend_id  IS '主键';
COMMENT ON COLUMN public.friend.name       IS '姓名';
COMMENT ON COLUMN public.friend.first_know IS '认识时间';
COMMENT ON COLUMN public.friend.remark     IS '备注';
COMMENT ON COLUMN public.friend.creator    IS '创建人';

-- example
INSERT INTO public.friend VALUES (DEFAULT, DEFAULT, '潘瑜妍', NULL, '传媒院')


--
-- 好友称呼表
--
CREATE TABLE if not exists public.greet
(
  ts          timestamptz DEFAULT now(),
  id          VARCHAR(40),
  name        VARCHAR(20),
  tail_letter CHAR,
  pet_name    VARCHAR(32), 
  remark      VARCHAR     DEFAULT NULL,
  creator     VARCHAR(40) DEFAULT 'njupter',
  PRIMARY KEY (id)
);

COMMENT ON TABLE public.greet              IS '好友清单';
COMMENT ON COLUMN public.greet.ts          IS '创建时间';
COMMENT ON COLUMN public.greet.id          IS '主键';
COMMENT ON COLUMN public.greet.name        IS '姓名';
COMMENT ON COLUMN public.greet.tail_letter IS '姓名尾字母';
COMMENT ON COLUMN public.greet.pet_name    IS '昵称';
COMMENT ON COLUMN public.greet.remark      IS '备注, 比如不再使用';
COMMENT ON COLUMN public.greet.creator     IS '创建人';

-- example
INSERT INTO "public"."greet" VALUES (DEFAULT, 'nj001', '潘瑜妍', 'Y', '潘导', NULL);
INSERT INTO "public"."greet" VALUES (DEFAULT, 'nj002', '潘瑜妍', 'Y', '千江大女神', '已过时');
INSERT INTO "public"."greet" VALUES (DEFAULT, 'nj003', '许波於', 'Y', '小於');


--
-- 好友联系方式表
--
CREATE TABLE if not exists public.contact
(
  ts timestamptz DEFAULT now(),
  id serial,

  friend_id   UUID        NOT NULL,
  name        VARCHAR(20) NOT NULL,
  call_type   VARCHAR(12) NULL,
  number      VARCHAR(32) NULL,
  get_time    date        NULL, 
  preference  BOOLEAN     DEFAULT false, 
  remark      VARCHAR     DEFAULT NULL, 
  creator     VARCHAR(20) DEFAULT 'njupter',
  PRIMARY KEY (id)
)with (oids = false);

COMMENT ON TABLE public.contact             IS '联系方式';
COMMENT ON COLUMN public.contact.ts         IS '创建时间';
COMMENT ON COLUMN public.contact.id         IS '主键';
COMMENT ON COLUMN public.contact.friend_id  IS '好友id';
COMMENT ON COLUMN public.contact.name       IS '名字';
COMMENT ON COLUMN public.contact.call_type  IS '联系类型, 可选, qq|vx|weibo|blog|iphone';
COMMENT ON COLUMN public.contact.number     IS '该类型的号码';
COMMENT ON COLUMN public.contact.get_time   IS '获得类型方式的时间';
COMMENT ON COLUMN public.contact.preference IS '首选方式, 只有一个是 true';
COMMENT ON COLUMN public.contact.remark     IS '备注';
COMMENT ON COLUMN public.contact.creator    IS '创建人';

INSERT INTO public.contact VALUES (DEFAULT, DEFAULT, 'xxx', '潘瑜妍', 'vx', '1', '2018-01-01', true);
