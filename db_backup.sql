-- user_table definition
DROP TABLE IF EXISTS `tracking_order`;
DROP  TABLE IF EXISTS `user_table`; 
CREATE TABLE `user_table` (
  `Id` varchar(36) NOT NULL,
  `Name` varchar(250) NOT NULL,
  `email` varchar(50) NOT NULL,
  `password` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `order` definition

DROP TABLE IF EXISTS `order`;
CREATE TABLE `order` (
  `id` varchar(36) NOT NULL,
  `sender` varchar(250) DEFAULT NULL,
  `sender_mobile_no` varchar(20) DEFAULT NULL,
  `receiver_name` varchar(250) DEFAULT NULL,
  `receiver_address` varchar(250) NOT NULL,
  `weight` int(11) NOT NULL,
  `status` int(11) NOT NULL,
  `receiver_mobile_no` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- `tracking_order` definition

CREATE TABLE `tracking_order` (
	id varchar(36) NOT NULL PRIMARY KEY,
	order_id varchar(36) NOT NULL,
	check_points text NULL,
	time_stamp datetime,
	CONSTRAINT fk_tracking_order_order FOREIGN KEY (order_id) REFERENCES `order`(id)
) 	ENGINE=InnoDB DEFAULT CHARSET=utf8;
