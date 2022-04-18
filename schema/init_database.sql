CREATE TABLE book (
    id         serial unique,
    name       varchar(255),
    category   varchar(255),
    content    varchar(255),
    author     varchar(255),
);

CREATE TABLE school (
    id serial unique,
    school_name varchar(255) not null,
    address varchar(255),
    landlines_number varchar(22),
    phone_number varchar(22),
    fax_number varchar(22),
    hot_line varchar(22),
    ward varchar(255),
    city varchar(255),
    international varchar(255),
    email varchar(255),
    logo varchar(255)
);

CREATE TABLE department (
    id              int unique,
    code            varchar(255),
    name            varchar(255),
    school_id       int,
    user_manager_id serial,
    avatar varchar(255)
);

CREATE TABLE class (
    id serial unique,
    code varchar(255),
    name varchar(255),
    subject_name varchar(255),
    note text,
    class_type_id serial,
    start_date date,
    end_date date,
    address varchar(255)
);

CREATE TABLE staff (
    id serial constraint staff_staff_id_key
            unique,
    avatar varchar(255),
    name varchar(255),
    family_name varchar(255),
    middle_name varchar(255),
    gender varchar(3),
    birth_day date,
    nationality varchar(100),
    department_id int,
    level int,
    qualification varchar(100),
    subject varchar(100)[],
    join_date date,
    experience varchar(255),
    current_address varchar(255),
    permanent_address varchar (255),
    phone_number varchar(100),
    email varchar(255),
    creation_date timestamp  default  NOW(),
    modification_date timestamp default NOW()
);

create table if not exists user_account
(
    id             serial not null
    constraint user_account_pkey
    primary key,
    user_name      varchar(255) unique,
    user_email     varchar(255) unique,
    user_password  varchar(255),
    is_active      boolean   default false,
    created_at     timestamp default now(),
    modified_at    timestamp default now(),
    modified_by    integer,
    created_by     integer,
    phone_number   varchar(255) unique,
    login_count    integer,
    is_first_login boolean,
    surname        varchar(255),
    first_name     varchar(255),
    avatar         text
);

-- TODO mapping user with teacher, security or student
CREATE TABLE user_info (
    id serial constraint user_user_id_key
            unique,
    user_account_id serial unique,
    school_id serial,
    department_id serial,
    class_id int[],
    role_id int,
    creation_date timestamp  default  NOW(),
    modification_date timestamp default NOW(),
    is_active boolean default  false,
    status int
);

CREATE TABLE role (
    id serial constraint role_pk unique ,
    title varchar(100),
    group_role_id int[],
    description text,
    creation_date timestamp  default  NOW(),
    modification_date timestamp default NOW(),
    status int
);

CREATE TABLE role_group (
    id serial unique,
    name varchar (255),
    group_role_action_id int[]
);

CREATE TABLE role_action (
    id serial unique,
    action_name varchar(255)
);

CREATE TABLE role_school (
    id serial unique,
    role_id serial,
    school_id serial
);

-- role 1, group_role_id [1,2,3]

-- group_role_id 1, group_role_action_id [1,2,3]

-- role action 1, "READ"
-- role action 2, "UPDATE"
-- role action 3, "DELETE"

INSERT INTO user_account VALUES (1, 'admin', 'sa@example.com', '123456', true)

INSERT INTO public.school (id, school_name, address, landlines_number, phone_number, fax_number, hot_line, ward, city, international, email, logo) VALUES (DEFAULT, 'LSH', 'Test', '23223', '23223', '232', '232', 'DN', 'DN', 'VN', 'school@example.com', 'http://')

INSERT INTO public.user_info (id, user_account_id, school_id, department_id, class_id, role_id, creation_date, modification_date, is_active, status) VALUES (DEFAULT, 1, 1, 1, '{1}', 1, '2021-06-28 13:08:57.000000', '2021-06-28 13:08:59.000000', DEFAULT, null)

INSERT INTO role_action VALUES (1, 'Read'), (2, 'Write'), (3, 'Update'), (4, 'Delete');

INSERT INTO role_group VALUES  (1, 'school', ARRAY [1, 2, 3, 4]);

INSERT INTO role VALUES (1, 'SUPER ADMIN', ARRAY [1], 'Role Full Quyền' , now(), now(), 1),
                        (2, 'Admin', ARRAY [1], 'Admin' , now(), now(), 1);
