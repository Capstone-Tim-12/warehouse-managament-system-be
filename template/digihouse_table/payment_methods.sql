-- -------------------------------------------------------------
-- TablePlus 5.6.2(516)
--
-- https://tableplus.com/
--
-- Database: fadilah_db
-- Generation Time: 2023-12-16 14:16:29.8640
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


INSERT INTO `payment_methods` (`id`, `name`, `image`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'ATM/Bank transfer', 'https://www.dropbox.com/scl/fi/itagl454kixu0pquoerb0/file_mG6hpX1hdDf3SA3xiJqC.png?rlkey=uyy6baz40li69rkihnux0jhza&raw=1', '2023-12-03 06:21:03.626', '2023-12-03 06:21:03.626', NULL),
(2, 'Kartu Kredit/Debit', 'https://www.dropbox.com/scl/fi/mtn1oer4c03mgayq696tv/file_Vw9AeIcop8LBqPVSGfvN.png?rlkey=vk2ifm1zozaymcmdbsdhlf6dm&raw=1', '2023-12-03 06:21:03.626', '2023-12-03 06:21:03.626', NULL);


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;