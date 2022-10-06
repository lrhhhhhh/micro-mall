create database if not exists user;
use user;
create table if not exists user (
    id              bigint unsigned auto_increment primary key,
    username        varchar(191) not null,
    password_digest longtext     not null,
    created_at      datetime(3)  null,
    updated_at      datetime(3)  null,
    deleted_at      datetime(3)  null,
    constraint username unique (username)
)engine=innodb, charset=utf8mb4;


create database if not exists stock;
use stock;
create table if not exists stock(
    `id`       int primary key not null auto_increment,
    `goods_id` int not null,
    `count`    int not null
)engine=innodb, charset=utf8mb4;

create table if not exists stock_task(
    id       bigint primary key not null auto_increment,
    stock_id int not null,
    amount   int not null
)engine=innodb, charset=utf8mb4, comment '顺序插入库存扣减流水，避免秒杀对同一行行锁的并发争用';


create database if not exists `order`;
use `order`;
create table if not exists `order` (
    id          bigint auto_increment primary key,
    uid         bigint not null,
    activity_id bigint not null,
    goods_id    bigint not null,
    stock_id    bigint not null,
    count       bigint not null,
    status      bigint not null default 0 comment '-2超时取消，-1回滚取消，1支付成功',
    created_at  bigint not null,
    updated_at  bigint not null,
    deleted_at  bigint not null
);


create database if not exists activity;
use activity;
create table if not exists activity(
    id              bigint unsigned auto_increment primary key,
    activity_name   longtext    not null,
    goods_id        bigint      not null,
    stock_id        bigint      not null,
    start_time      bigint      not null,
    end_time        bigint      not null,
    total           bigint      not null,
    status          int         not null,
    buy_limit       bigint      not null,
    buy_probability double      not null
);





create database if not exists dtm_barrier
    /*!40100 DEFAULT CHARACTER SET utf8mb4 */
;
drop table if exists dtm_barrier.barrier;
create table if not exists dtm_barrier.barrier(
    id bigint(22) PRIMARY KEY AUTO_INCREMENT,
    trans_type varchar(45) default '',
    gid varchar(128) default '',
    branch_id varchar(128) default '',
    op varchar(45) default '',
    barrier_id varchar(45) default '',
    reason varchar(45) default '' comment 'the branch type who insert this record',
    create_time datetime DEFAULT now(),
    update_time datetime DEFAULT now(),
    key(create_time),
    key(update_time),
    UNIQUE key(gid, branch_id, op, barrier_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;




CREATE DATABASE IF NOT EXISTS dtm
    /*!40100 DEFAULT CHARACTER SET utf8mb4 */
;

CREATE TABLE if not EXISTS dtm.trans_global (
    `id` bigint(22) NOT NULL AUTO_INCREMENT,
    `gid` varchar(128) NOT NULL COMMENT 'global transaction id',
    `trans_type` varchar(45) not null COMMENT 'transaction type: saga | xa | tcc | msg',
    `status` varchar(12) NOT NULL COMMENT 'tranaction status: prepared | submitted | aborting | finished | rollbacked',
    `query_prepared` varchar(1024) NOT NULL COMMENT 'url to check for msg|workflow',
    `protocol` varchar(45) not null comment 'protocol: http | grpc | json-rpc',
    `create_time` datetime DEFAULT NULL,
    `update_time` datetime DEFAULT NULL,
    `finish_time` datetime DEFAULT NULL,
    `rollback_time` datetime DEFAULT NULL,
    `options` varchar(1024) DEFAULT 'options for transaction like: TimeoutToFail, RequestTimeout',
    `custom_data` varchar(1024) DEFAULT '' COMMENT 'custom data for transaction',
    `next_cron_interval` int(11) default null comment 'next cron interval. for use of cron job',
    `next_cron_time` datetime default null comment 'next time to process this trans. for use of cron job',
    `owner` varchar(128) not null default '' comment 'who is locking this trans',
    `ext_data` TEXT comment 'result for this trans. currently used in workflow pattern',
    `result` varchar(1024) DEFAULT '' COMMENT 'rollback reason for transaction',
    `rollback_reason` varchar(1024) DEFAULT '' COMMENT 'rollback reason for transaction',
    PRIMARY KEY (`id`),
    UNIQUE KEY `gid` (`gid`),
    key `owner`(`owner`),
    key `status_next_cron_time` (`status`, `next_cron_time`) comment 'cron job will use this index to query trans'
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS dtm.trans_branch_op (
    `id` bigint(22) NOT NULL AUTO_INCREMENT,
    `gid` varchar(128) NOT NULL COMMENT 'global transaction id',
    `url` varchar(1024) NOT NULL COMMENT 'the url of this op',
    `data` TEXT COMMENT 'request body, depreceated',
    `bin_data` BLOB COMMENT 'request body',
    `branch_id` VARCHAR(128) NOT NULL COMMENT 'transaction branch ID',
    `op` varchar(45) NOT NULL COMMENT 'transaction operation type like: action | compensate | try | confirm | cancel',
    `status` varchar(45) NOT NULL COMMENT 'transaction op status: prepared | succeed | failed',
    `finish_time` datetime DEFAULT NULL,
    `rollback_time` datetime DEFAULT NULL,
    `create_time` datetime DEFAULT NULL,
    `update_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `gid_uniq` (`gid`, `branch_id`, `op`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


CREATE TABLE IF NOT EXISTS dtm.kv (
    `id` bigint(22) NOT NULL AUTO_INCREMENT,
    `cat` varchar(45) NOT NULL COMMENT 'the category of this data',
    `k` varchar(128) NOT NULL,
    `v` TEXT,
    `version` bigint(22) default 1 COMMENT 'version of the value',
    create_time datetime default NULL,
    update_time datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE key `uniq_k`(`cat`, `k`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
