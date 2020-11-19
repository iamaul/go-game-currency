-- phpMyAdmin SQL Dump
-- version 5.0.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Nov 19, 2020 at 07:46 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.9

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `game_currency_betest_whim`
--
CREATE DATABASE IF NOT EXISTS `game_currency_betest_whim` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `game_currency_betest_whim`;

-- --------------------------------------------------------

--
-- Table structure for table `conversion_currencies`
--

CREATE TABLE `conversion_currencies` (
  `id` int(11) NOT NULL,
  `currency_id_from` int(11) NOT NULL,
  `currency_id_to` int(11) NOT NULL,
  `amount` bigint(11) NOT NULL,
  `result` bigint(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `conversion_currencies`
--

INSERT INTO `conversion_currencies` (`id`, `currency_id_from`, `currency_id_to`, `amount`, `result`, `created_at`, `updated_at`) VALUES
(1, 1, 2, 580, 20, '2020-11-19 15:00:42', NULL),
(2, 2, 1, 20, 580, '2020-11-19 15:01:24', NULL),
(3, 3, 1, 20, 9860, '2020-11-19 15:03:29', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `conversion_rates`
--

CREATE TABLE `conversion_rates` (
  `id` int(11) NOT NULL,
  `currency_id_from` int(11) NOT NULL,
  `currency_id_to` int(11) NOT NULL,
  `rate` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `conversion_rates`
--

INSERT INTO `conversion_rates` (`id`, `currency_id_from`, `currency_id_to`, `rate`, `created_at`, `updated_at`) VALUES
(1, 2, 1, 29, '2020-11-19 12:28:01', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `currencies`
--

CREATE TABLE `currencies` (
  `id` int(11) NOT NULL,
  `name` varchar(64) NOT NULL,
  `created_at` timestamp NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `currencies`
--

INSERT INTO `currencies` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Knut', '2020-11-19 06:39:45', NULL),
(2, 'Sickle', '2020-11-19 06:39:53', NULL),
(3, 'Galleon', '2020-11-19 06:39:59', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `conversion_currencies`
--
ALTER TABLE `conversion_currencies`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `conversion_rates`
--
ALTER TABLE `conversion_rates`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `currencies`
--
ALTER TABLE `currencies`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `conversion_currencies`
--
ALTER TABLE `conversion_currencies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `conversion_rates`
--
ALTER TABLE `conversion_rates`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `currencies`
--
ALTER TABLE `currencies`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
