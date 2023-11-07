create table province(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255)
);

create table regency(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    province_id int(11),
    name varchar(255),
    FOREIGN KEY (province_id) REFERENCES province(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table district(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    regency_id int(11),
    name varchar(255),
    FOREIGN KEY (regency_id) REFERENCES regency(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table village(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    district_id int(11),
    name varchar(255),
    FOREIGN KEY (district_id) REFERENCES district(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table `user`(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username varchar(35),
    email varchar(100),
    is_verify_acount bool,
    is_verify_identity bool,
    password text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table user_detail(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	photo varchar(255),
    nik varchar(16),
    full_name varchar(100),
    gender varchar(200),
    place_of_bird date,
    works varchar(100),
    citizenship varchar(100),
    user_id int(11),
    province_id int(11),
    regency_id int(11),
    district_id int(11),
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
    FOREIGN KEY (user_id) REFERENCES `user`(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (province_id) REFERENCES province(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (regency_id) REFERENCES regency(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (district_id) REFERENCES district(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table warehouse(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    name varchar(255),
    province_id int(10),
    regency_id int(10),
    wide varchar(200),
    owner varchar(100),
    phone_number varchar(15),
    price bigint,
    status enum('tesedia', 'disewa', 'dalam pemeliharaan'),
    description text,
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
	FOREIGN KEY (province_id) REFERENCES province(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (regency_id) REFERENCES regency(id) ON DELETE CASCADE ON UPDATE CASCADE
);

create table capacity(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
	created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL
);

create table warehouse_picture(
	id int PRIMARY KEY NOT NULL AUTO_INCREMENT,
    picture varchar(255),
    warehouse_id int(12),
    created_at datetime(3) DEFAULT NULL,
	updated_at datetime(3) DEFAULT NULL,
	deleted_at datetime(3) DEFAULT NULL,
    FOREIGN KEY (warehouse_id) REFERENCES warehouse(id) ON DELETE CASCADE ON UPDATE CASCADE
);