-- MySQL dump 10.13  Distrib 8.0.45, for Linux (x86_64)
--
-- Host: localhost    Database: subscribers
-- ------------------------------------------------------
-- Server version	8.0.45-0ubuntu0.24.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `clients` (
  `id` int NOT NULL AUTO_INCREMENT,
  `fio` varchar(255) DEFAULT NULL,
  `city` varchar(100) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `phone` varchar(50) DEFAULT NULL,
  `ip` varchar(50) DEFAULT NULL,
  `login` varchar(100) DEFAULT NULL,
  `password` varchar(100) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `clients`
--

LOCK TABLES `clients` WRITE;
/*!40000 ALTER TABLE `clients` DISABLE KEYS */;
INSERT INTO `clients` VALUES (2,'Лиана Мармонт','Медвежий остров','Остров','+77786578943','10.88.11.100','l.marmont@gmail.com',NULL,'2026-03-18 15:53:12'),(3,'Нед Старк','Винтерфел','Крипта','+77022559784','10.88.11.101','ned.stark@gmail.com',NULL,'2026-03-18 15:54:08'),(4,'Джейме Ланистер','Королевская гавань','Королевский замок','+77028558512','10.88.11.102','dj.lanister@gmail.com',NULL,'2026-03-18 15:55:34'),(5,'Артур Дэйн','Звездопад','Замок ','+77772456891','10.88.11.103','a.dayn@gmail.com',NULL,'2026-03-18 16:14:28'),(6,'Станис Баратеон','Драконий камень','Замок','+77788965471','10.88.11.104','s.barateon@gmail.com',NULL,'2026-03-18 16:16:13'),(8,'Дайнерис Таргариен','Дотрак ','Дотракийская степь','+77785964872','10.88.11.105','d.targaryen@gmail.com',NULL,'2026-03-18 16:20:59'),(9,'Эдмонт Талли','Риверан','Замок','+77785964845','10.88.11.106','a.talli@gmail.com',NULL,'2026-03-18 16:22:23'),(10,'Джон Амбер','Последний очаг','Замок','+77785978931','10.88.11.107','dj.amber@gmail.com',NULL,'2026-03-18 16:24:55'),(11,'Лиза Аррен','Орлиное гнездо','Замок','+77785978912','10.88.11.108','l.arren@gmail.com',NULL,'2026-03-18 16:26:15'),(12,'Бэйлон Грэйджой','Железные острова','Остров','+77785975417','10.88.11.109','b.greydjoy@gmail.com',NULL,'2026-03-18 16:29:26');
/*!40000 ALTER TABLE `clients` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-03-25 11:21:28
