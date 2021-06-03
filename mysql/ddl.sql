create table student
(
	id varchar(20) primary key,
    name varchar(20),
    password varchar(30) not null
);

create table teacher
(
	id varchar(20) primary key,
    name varchar(20),
    password varchar(30) not null
);

create table class
(
    id varchar(40) primary key,
    name varchar(20),
    teacher_id varchar(20),
    picURL varchar(200),
    foreign key (teacher_id) references teacher(id)
);

create table student_class
(
    student_id varchar(20) not null,
    class_id varchar(40) not null,
    foreign key (student_id) references student(id),
    foreign key (class_id) references class(id),
    primary key (student_id, class_id)
);

create table homework_info
(
	pid varchar(40) primary key,
    name varchar(20),
    description text,
    start_time date,
    deadline date,
    class_id varchar(40),
    foreign key (class_id) references class(id)
);

create table homework
(
    pid varchar(40),
    student_id varchar(20),
    class_id varchar(40),
    score float,
    commit text,
    url varchar(200),
    foreign key (pid) references homework_info(pid),
    foreign key (student_id) references student(id),
    foreign key (class_id) references class(id),
    primary key (pid, student_id, class_id)
);