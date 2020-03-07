-- +migrate Up

---------- 应用表 ----------
CREATE TABLE `app`(
  id         integer     not null primary key autoincrement,
  app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
  app_key    VARCHAR(40)   not null DEFAULT '', -- 应用key
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
)


-- +migrate StatementBegin
CREATE TRIGGER app_updated_at
  AFTER UPDATE
  ON `app`
  BEGIN
    update `app` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd

---------- 访客信息表 ----------
CREATE TABLE `visitor`
(
  id         integer     not null primary key autoincrement,
  app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
  sid        VARCHAR(40) not null default '', -- 访客唯一ID
  browser    VARCHAR(255) not null default '', -- 浏览器
  os         VARCHAR(255) not null default '', -- 操作系统
  ip         VARCHAR(20) not null default '', -- IP地址
  address    VARCHAR(20) not null default '', -- 访问地点
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
);

-- +migrate StatementBegin
CREATE TRIGGER visitor_updated_at
  AFTER UPDATE
  ON `visitor`
  BEGIN
    update `visitor` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd

---------- 频道表 ----------

CREATE TABLE `channel`
(
  id         integer       not null primary key autoincrement,
  app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
  channel_id   VARCHAR(40) not null default '', -- 频道ID
  channel_type smallint    not null default 0,  -- 频道类型 1.单聊 2.群聊 3.机器人
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
);

-- +migrate StatementBegin
CREATE TRIGGER channel_updated_at
  AFTER UPDATE
  ON `channel`
  BEGIN
    update `channel` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd

---------- 订阅者表 ----------
create table `subscriber`
(
  id         integer     not null primary key autoincrement,
  app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
  channel_id VARCHAR(40) not null default '', -- 频道ID
  channel_type smallint    not null default 0,  -- 频道类型 1.单聊 2.群聊 3.机器人
  uid        VARCHAR(40) not null default '', -- 订阅用户的uid
  created_at BIGINT      not null default 0,  -- 创建时间
  updated_at BIGINT      not null default 0   -- 更新时间
);

---------- 路由历史表 ----------

CREATE TABLE `route_history`(
  id           integer       not null primary key autoincrement,
  app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
  visitor_sid  VARCHAR(40)   not null default '',  -- 访客sid
  channel_id VARCHAR(40) not null default '', -- 接入的频道ID
  channel_type smallint    not null default 0,  -- 接入的频道类型
  status        smallint    not NULL default  0,   --  1.接入成功 2.接入失败
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
)

-- +migrate StatementBegin
CREATE TRIGGER route_history_updated_at
  AFTER UPDATE
  ON `route_history`
  BEGIN
    update `route_history` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd


---------- 消息表 ----------
create table `message`
(
  id           integer          not null primary key autoincrement,
  app_id        VARCHAR(40)   not null DEFAULT '', -- 应用ID
  message_id   UNSIGNED BIG INT not null default 0,  -- 消息唯一ID（全局唯一）
  message_seq  UNSIGNED BIG INT not null default 0,  -- 消息序列号(非严格递增)
  from_uid     VARCHAR(40)      not null default '', -- 发送者uid
  to_uid       VARCHAR(40)      not null default '', -- 接收者uid
  channel_id   VARCHAR(40)      not null default '', -- 频道ID
  channel_type smallint         not null default 0,  -- 频道类型
  timestamp    BIGINT           not null default 0,  -- 消息时间
  payload      blob             not null default '', -- 消息内容
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
);
-- +migrate StatementBegin
CREATE TRIGGER message_updated_at
  AFTER UPDATE
  ON `message`
  BEGIN
    update `message` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd

---------- 最近会话表 ----------

CREATE TABLE `conversation`(
    id           integer       not null primary key autoincrement,
    app_id       VARCHAR(40)   not null DEFAULT '', -- 应用ID
    channel_id   VARCHAR(40) not null default '', -- 频道ID
    channel_type smallint    not null default 0,  -- 频道类型
    last_msg_id  bigint      not null default 0, -- 最后一条消息ID
    last_msg_time integer    not null default 0, -- 最后一次消息时间戳（长度为10为的时间戳到秒）
    created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
    updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
)
-- +migrate StatementBegin
CREATE TRIGGER conversation_updated_at
  AFTER UPDATE
  ON `conversation`
  BEGIN
    update `conversation` SET updated_at = datetime('now') WHERE id = NEW.id;
  END;
-- +migrate StatementEnd

--------- 技能表 ---------

CREATE TABLE `skill`(
    id  integer  not null primary key autoincrement,
    app_id     VARCHAR(40)   not null DEFAULT '', -- 应用ID
    skill_no    VARCHAR(40)   not null DEFAULT '', -- 技能组编号
    name  VARCHAR(40) not null default '', -- 名称
    remember smallint not null default 0, -- 是否启用记忆分配
    strategy integer  not null default 0, --分配策略 1. 优先 2.最少  3. 随机 4.轮流
    created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
    updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
);


--------- 技能与频道的关系表 ---------

CREATE TABLE `skill_channel`(
  id          integer      not null primary key autoincrement,
  app_id      VARCHAR(40)   not null DEFAULT '', -- 应用ID
  skill_no    VARCHAR(40)   not null default '',  -- 技能编号
  channel_id   VARCHAR(40) not null default '', -- 频道ID （客服的频道ID）
  channel_type smallint    not null default 0,  -- 频道类型
  priority     integer        not null default 0, -- 数字越小，级别越高，优先级相同的采用轮询策略，默认为随机策略
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间

);

--------- 聊天窗口样式表 ---------

CREATE TABLE `chatpopup_setting`(
  id          integer      not null primary key autoincrement,
  app_id      VARCHAR(40)   not null DEFAULT '', -- 应用ID
  chatpopup_no VARCHAR(40)  not null DEFAULT '', -- 聊天窗口唯一编号
  site_name   VARCHAR(40)   not null DEFAULT '', -- 站点名称
  site_url    VARCHAR(255)  not null DEFAULT '', -- 站点域名
  site_title  VARCHAR(100)   not null DEFAULT '', -- 站点标题
  company     VARCHAR(100)  not null default '', -- 公司名称
  company_logo VARCHAR(255) not null default '', -- 公司logo
  show_nickname smallint    not null default  0, -- 是否显示昵称
  avatar       VARCHAR(255) not null default '', -- 系统头像
  welcomme    VARCHAR(255)  not null default '', -- 系统欢迎语
  show_forward_people smallint  not null default 0, -- 是否显示转人工 0.是否 1.是
  robot_no    VARCHAR(40) not null default '', -- 机器人编号，如果存在则将自动接入机器人             
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
)

--------- 聊天窗口与技能的关系 ---------

CREATE TABLE `chatpopup_skill`(
  id          integer      not null primary key autoincrement,
  app_id      VARCHAR(40)   not null DEFAULT '', -- 应用ID
  chatpopup_no VARCHAR(40)  not null DEFAULT '', -- 聊天窗口唯一编号
  skill_no    VARCHAR(40)   not null DEFAULT '', -- 技能编号
  created_at timeStamp     not null DEFAULT (datetime('now', 'localtime')), -- 创建时间
  updated_at timeStamp     not null DEFAULT (datetime('now', 'localtime'))  -- 更新时间
)

--------- 通话队列 ---------


## 消息通知框是否弹出设置

## 禁音设置