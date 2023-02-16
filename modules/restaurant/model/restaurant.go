package model

//CREATE TABLE `restaurants` (
//  `id` int(11) NOT NULL AUTO_INCREMENT,
//  `owner_id` int(11) DEFAULT NULL,
//  `name` varchar(50) NOT NULL,
//  `addr` varchar(255) NOT NULL,
//  `city_id` int(11) DEFAULT NULL,
//  `lat` double DEFAULT NULL,
//  `lng` double DEFAULT NULL,
//  `cover` json DEFAULT NULL,
//  `logo` json DEFAULT NULL,
//  `shipping_fee_per_km` double DEFAULT '0',
//  `status` int(11) NOT NULL DEFAULT '1',
//  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
//  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
//  PRIMARY KEY (`id`),
//  KEY `owner_id` (`owner_id`) USING BTREE,
//  KEY `city_id` (`city_id`) USING BTREE,
//  KEY `status` (`status`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8;

type Restaurant struct {
	Id   int    `json:"id,omitempty" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
