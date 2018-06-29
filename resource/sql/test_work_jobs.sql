CREATE DATABASE  IF NOT EXISTS `test_work` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `test_work`;
-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: localhost    Database: test_work
-- ------------------------------------------------------
-- Server version	5.7.21-0ubuntu0.16.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `jobs`
--

DROP TABLE IF EXISTS `jobs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `jobs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company` varchar(45) NOT NULL,
  `boss` varchar(45) NOT NULL,
  `city` varchar(45) NOT NULL,
  `address` varchar(45) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `jobs`
--

LOCK TABLES `jobs` WRITE;
/*!40000 ALTER TABLE `jobs` DISABLE KEYS */;
INSERT INTO `jobs` VALUES (5,'alijiujiu','mayun','hangzhou','shi-zhong-xin'),(11,'alibaba','mayun','hangzhou','shi-zhong-xin'),(12,'alibaba','mayun','hangzhou','shi-zhong-xin'),(13,'alibaba','mayun','hangzhou','shi-zhong-xin'),(15,'baidu','liyanhong','beijing','xi-er-qi'),(16,'tengxun','mahuateng','beijing','wu-dao-kou'),(17,'alibaba','mayun','hangzhou','shi-zhong-xin'),(18,'baidu','liyanhong','beijing','xi-er-qi'),(19,'tengxun','mahuateng','beijing','wu-dao-kou'),(21,'baidu','liyanhong','beijing','xi-er-qi'),(22,'tengxun','mahuateng','beijing','wu-dao-kou'),(23,'alibaba','mayun','hangzhou','shi-zhong-xin'),(24,'baidu','liyanhong','beijing','xi-er-qi'),(25,'tengxun','mahuateng','beijing','wu-dao-kou'),(26,'alibaba','mayun','hangzhou','shi-zhong-xin'),(27,'baidu','liyanhong','beijing','xi-er-qi'),(28,'tengxun','mahuateng','beijing','wu-dao-kou'),(29,'alibaba','mayun','hangzhou','shi-zhong-xin');
/*!40000 ALTER TABLE `jobs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-06-25 13:47:17
