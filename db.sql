create table if not exists metrics
(
    name varchar(128) not null
    constraint metrics_pk
    primary key,
    labels jsonb not null,
    value numeric not null,
    description varchar(256),
    metric_time timestamp with time zone not null,
    create_time timestamp with time zone not null
                              );

comment on table metrics is '监控指标';

alter table metrics owner to root;

create table if not exists spans
(
    id varchar(256) not null
    constraint spans_pk
    primary key,
    parent_id varchar(256) not null,
    trace_id varchar(256) not null,
    service varchar(128) not null,
    operation varchar(256) not null,
    tags jsonb not null,
    logs jsonb not null,
    duration bigint not null,
    status integer default 0 not null,
    start_time timestamp with time zone not null,
    end_time timestamp with time zone not null,
    create_time timestamp with time zone not null
                              );

comment on table spans is '链路单元';

alter table spans owner to root;

