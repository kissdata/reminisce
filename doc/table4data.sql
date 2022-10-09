-- 建表语句示例
-- 存储交流数据

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

--
-- 最新叙旧表
--
CREATE TABLE if not exists chat_latest (
    ts         timestamptz,
    friend_id  UUID,
    pet_name   VARCHAR(20), 
    datetime   date, 
    chat_topic varchar(128), 
    content    text,
    remark     VARCHAR      DEFAULT NULL,
    creator    VARCHAR(20)  DEFAULT 'njupter',
    PRIMARY KEY (friend_id)
);

COMMENT ON TABLE public.chat_latest             IS '最新叙旧';
COMMENT ON COLUMN public.chat_latest.ts         IS '创建时间';
COMMENT ON COLUMN public.chat_latest.friend_id  IS '好友id';
COMMENT ON COLUMN public.chat_latest.pet_name   IS '好友昵称';
COMMENT ON COLUMN public.chat_latest.datetime   IS '交流当日';
COMMENT ON COLUMN public.chat_latest.chat_topic IS '交流话题';
COMMENT ON COLUMN public.chat_latest.content    IS '话题内容';
COMMENT ON COLUMN public.chat_latest.remark     IS '备注';
COMMENT ON COLUMN public.chat_latest.creator    IS '记录人';


--
-- 叙旧历史记录表
--
CREATE TABLE if not exists chat_history (
    ts         timestamptz,
    id         INTEGER GENERATED ALWAYS AS IDENTITY (cache 1), 
    friend_id  UUID,
    pet_name   VARCHAR(20), 
    datetime   date, 
    chat_topic varchar(128),
    content    text,
    remark     VARCHAR      DEFAULT NULL,
    creator    VARCHAR(20)  DEFAULT 'njupter',
    PRIMARY KEY (id)
);

COMMENT ON TABLE public.chat_history             IS '叙旧历史';
COMMENT ON COLUMN public.chat_history.ts         IS '创建时间';
COMMENT ON COLUMN public.chat_history.id         IS '主键';
COMMENT ON COLUMN public.chat_history.friend_id  IS '好友id';
COMMENT ON COLUMN public.chat_history.pet_name   IS '好友昵称';
COMMENT ON COLUMN public.chat_history.datetime   IS '交流当日';
COMMENT ON COLUMN public.chat_history.chat_topic IS '交流话题';
COMMENT ON COLUMN public.chat_history.content    IS '话题内容';
COMMENT ON COLUMN public.chat_history.remark     IS '备注';
COMMENT ON COLUMN public.chat_history.creator    IS '记录人';

-- Tips
-- don't use serial