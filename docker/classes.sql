SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `classes`
--
DROP DATABASE IF EXISTS classes;
CREATE DATABASE IF NOT EXISTS `classes` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `classes`;



DROP TABLE IF EXISTS classes;
CREATE TABLE IF NOT EXISTS classes (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  duration INT,
  date_time DATETIME,
  suitable_level VARCHAR(255),
  capacity INT,
  max_capacity INT,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



INSERT INTO classes (id, name, description, duration, date_time, suitable_level, capacity, max_capacity) VALUES
(1, 'Yoga', 'Relaxing and rejuvenating yoga class for all levels.', 60, '2022-03-18 09:00:00', 'All Levels', 30, 30),
(2, 'Pilates', 'Core-strengthening Pilates class focusing on flexibility and strength.', 45, '2022-03-19 17:30:00', 'Intermediate', 20, 20),
(3, 'Zumba', 'High-energy Zumba class with Latin-inspired dance moves.', 60, '2022-03-20 18:00:00', 'All Levels', 30, 30),
(4, 'Spin', 'Intense indoor cycling class for a great cardio workout.', 45, '2022-03-21 19:00:00', 'Advanced', 20, 20),
(5, 'HIIT', 'High-intensity interval training for a full-body workout.', 30, '2022-03-22 18:30:00', 'Intermediate', 25, 25);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;