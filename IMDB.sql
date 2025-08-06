/*
 Navicat Premium Data Transfer

 Source Server         : IMDB
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : imdb

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 20/04/2022 02:06:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cast
-- ----------------------------
DROP TABLE IF EXISTS `cast`;
CREATE TABLE `cast`  (
  `cast_id` int NOT NULL AUTO_INCREMENT,
  `cast_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cast_description` varchar(1000) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `cast_image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`cast_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cast
-- ----------------------------
INSERT INTO `cast` VALUES (1, 'Tao', 'master of actors', 'l');
INSERT INTO `cast` VALUES (2, 'zhang', 'PHD', NULL);
INSERT INTO `cast` VALUES (3, 'test', 'She is a good actress.', 'Tom.png');
INSERT INTO `cast` VALUES (4, 'Leonardo DiCaprio', 'One of the most famous actor', 'leonardo.png');
INSERT INTO `cast` VALUES (5, 'Brie Larson', 'Brie Larson has built an impressive career as an acclaimed television actress, rising feature film star and emerging recording artist.', 'brie-larson.png');
INSERT INTO `cast` VALUES (6, 'Joaquin Phoenix', 'Joaquin Phoenix was born Joaquin Rafael Bottom in San Juan, Puerto Rico, to Arlyn (Dunetz) and John Bottom, and is the middle child in a brood of five', 'Joaquin.jpg');
INSERT INTO `cast` VALUES (7, 'Tim Robbins', 'Born in West Covina, California, but raised in New York City, Tim Robbins is the son of former The Highwaymen singer Gil Robbins and actress Mary Robbins (née Bledsoe)', 'Tim.jpg');
INSERT INTO `cast` VALUES (8, 'Tom Hanks', 'OThomas Jeffrey Hanks was born in Concord, California, to Janet Marylyn (Frager), a hospital worker, and Amos Mefford Hanks, an itinerant cook.', 'tom.jpg');
INSERT INTO `cast` VALUES (9, 'Chadwick Boseman', 'Chadwick Boseman was an American actor. He is known for his portrayal of TChalla  Black Panther in the Marvel Cinematic Universe from 2016 to 2019, particularly in Black Panther (2018), and for his starring roles as several pioneering Americans, Jackie Robinson in 42 (2013).', 'chad.jpg');
INSERT INTO `cast` VALUES (10, 'Jamie Foxx', 'Jamie Foxx is an American actor, singer and comedian. He won an Academy Award for Best Actor, BAFTA Award for Best Actor in a Leading Role, and Golden Globe Award for Best Actor in a Musical or Comedy, for his work in the biographical film Ray (2004)', 'jamie.jpg');
INSERT INTO `cast` VALUES (11, 'Emilia Jones', 'Emilia is best known for playing the lead role of Ruby Rossi in CODA which premiered at Sundance Film Festival in 2021 and won more awards than any film in Sundance history', 'emilia.jpg');
INSERT INTO `cast` VALUES (12, 'Matthew McConaughey', 'American actor and producer Matthew David McConaughey was born in Uvalde, Texas', 'Matthew.jpg');
INSERT INTO `cast` VALUES (13, 'Uma Thurman', 'Uma Karuna Thurman was born in Boston, Massachusetts, into a highly unorthodox and internationally-minded family.', 'Uma.jpg');
INSERT INTO `cast` VALUES (14, 'Liam Neeson', 'Liam Neeson was born on June 7, 1952 in Ballymena, Northern Ireland, to Katherine (Brown), a cook, and Bernard Neeson, a school caretaker. He was raised in a Catholic household.', 'liam.jpg');
INSERT INTO `cast` VALUES (15, 'Tom Holland', 'Thomas Stanley Holland was born in Kingston-upon-Thames, Surrey, to Nicola Elizabeth (Frost), a photographer, and Dominic Holland (Dominic Anthony Holland), who is a comedian and author.', 'TomHolland.jpg');
INSERT INTO `cast` VALUES (16, 'Christian Bale', 'Christian Charles Philip Bale was born in Pembrokeshire, Wales, UK on January 30, 1974, to English parents Jennifer \"Jenny\" (James) and David Bale', 'bale.jpg');
INSERT INTO `cast` VALUES (17, 'Marlon Brando', 'Marlon Brando is widely considered the greatest movie actor of all time, rivaled only by the more theatrically oriented Laurence Olivier in terms of esteem.', 'marlon.jpg');
INSERT INTO `cast` VALUES (18, 'Elijah Wood', 'Elijah Wood is an American actor best known for portraying Frodo Baggins in Peter Jackson blockbuster Lord of the Rings film trilogy', 'Elijah.jpg');
INSERT INTO `cast` VALUES (19, 'Chris Hemsworth', 'Christopher Hemsworth was born on August 11, 1983 in Melbourne, Victoria, Australia to Leonie Hemsworth (née van Os), an English teacher & Craig Hemsworth, a social-services counselor.', 'Chris.jpg');

-- ----------------------------
-- Table structure for genre
-- ----------------------------
DROP TABLE IF EXISTS `genre`;
CREATE TABLE `genre`  (
  `genre_id` int NOT NULL AUTO_INCREMENT,
  `genre` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`genre_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of genre
-- ----------------------------
INSERT INTO `genre` VALUES (1, 'Comedy');
INSERT INTO `genre` VALUES (2, 'Thriller');
INSERT INTO `genre` VALUES (3, 'Action');
INSERT INTO `genre` VALUES (4, 'Scientific-fiction');
INSERT INTO `genre` VALUES (5, 'Adventure');
INSERT INTO `genre` VALUES (6, 'Romantic');
INSERT INTO `genre` VALUES (7, 'Documentary');
INSERT INTO `genre` VALUES (8, 'Drama');

-- ----------------------------
-- Table structure for movie
-- ----------------------------
DROP TABLE IF EXISTS `movie`;
CREATE TABLE `movie`  (
  `movie_id` int NOT NULL AUTO_INCREMENT,
  `movie_name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `year` int NULL DEFAULT NULL,
  `grade` float NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `image` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`movie_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie
-- ----------------------------
INSERT INTO `movie` VALUES (1, 'Avengers-Endgame', 2019, 5, 'After the devastating events of Avengers: Infinity War (2018), the universe is in ruins. With the help of remaining allies, the Avengers assemble once more in order to reverse Thanos\' actions and restore balance to the universe.', 'Avengers-Endgame.jpg');
INSERT INTO `movie` VALUES (2, 'Joker', 2019, 5, 'A mentally troubled stand-up comedian embarks on a downward spiral that leads to the creation of an iconic villain.', 'Joker.jpg');
INSERT INTO `movie` VALUES (3, 'Black Panther', 2018, 73, 'T\'Challa, heir to the hidden but advanced kingdom of Wakanda, must step forward to lead his people into a new future and must confront a challenger from his country\'s past.', 'black-panther.jpg');
INSERT INTO `movie` VALUES (4, 'test', 2001, 98, 'awesome!', 'logo.png');
INSERT INTO `movie` VALUES (5, 'test1', 2222, 4.5, '111111', 'Avengers-Endgame.jpg');
INSERT INTO `movie` VALUES (6, 'Captain Marvel', 2019, 68, 'Carol Danvers becomes one of the universe\'s most powerful heroes when Earth is caught in the middle of a galactic war between two alien races.', 'Captain-Marvel.jpg');
INSERT INTO `movie` VALUES (7, '\r\nDjango Unchained', 2012, 84, 'With the help of a German bounty-hunter, a freed slave sets out to rescue his wife from a brutal plantation-owner in Mississippi.', 'Django.jpg');
INSERT INTO `movie` VALUES (8, 'CODA', 2021, 80, 'As a CODA (Child of Deaf Adults) Ruby is the only hearing person in her deaf family. When the family\'s fishing business is threatened, Ruby finds herself torn between pursuing her passion at Berklee College of Music and her fear of abandoning her parents.', 'CODA.jpg');
INSERT INTO `movie` VALUES (9, 'Forrest Gump', 1984, NULL, 'The presidencies of Kennedy and Johnson, the Vietnam War, the Watergate scandal and other historical events unfold from the perspective of an Alabama man with an IQ of 75, whose only desire is to be reunited with his childhood sweetheart.', 'forrest.jpg');
INSERT INTO `movie` VALUES (10, 'Inception', 2012, 88, 'A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O., but his tragic past may doom the project and his team to disaster.', 'inception.jpg');
INSERT INTO `movie` VALUES (11, 'Interstellar', 2014, 86, 'A team of explorers travel through a wormhole in space in an attempt to ensure humanity\'s survival.', 'interstellar.jpg');
INSERT INTO `movie` VALUES (12, 'Pulp Fiction', 1994, 89, 'The lives of two mob hitmen, a boxer, a gangster and his wife, and a pair of diner bandits intertwine in four tales of violence and redemption.', 'Pulp-Fiction1.jpg');
INSERT INTO `movie` VALUES (13, 'Schindler\'s List', 1993, 90, 'In German-occupied Poland during World War II, industrialist Oskar Schindler gradually becomes concerned for his Jewish workforce after witnessing their persecution by the Nazis.', 'SchindlerList.jpg');
INSERT INTO `movie` VALUES (14, 'Spider-Man: No Way Home', 2021, NULL, 'With Spider-Man\'s identity now revealed, Peter asks Doctor Strange for help. When a spell goes wrong, dangerous foes from other worlds start to appear, forcing Peter to discover what it truly means to be Spider-Man.', 'Spider-Man3.jpg');
INSERT INTO `movie` VALUES (15, 'The Dark Knight', 2008, 90, 'When the menace known as the Joker wreaks havoc and chaos on the people of Gotham, Batman must accept one of the greatest psychological and physical tests of his ability to fight injustice.', 'TheDarkKnight.jpg');
INSERT INTO `movie` VALUES (16, 'The Godfather', 1972, 92, 'The aging patriarch of an organized crime dynasty in postwar New York City transfers control of his clandestine empire to his reluctant youngest son.', 'TheGodfather.jpg');
INSERT INTO `movie` VALUES (17, 'The Lord of the Rings: The Return of the King', 2003, 90, 'Gandalf and Aragorn lead the World of Men against Sauron\'s army to draw his gaze from Frodo and Sam as they approach Mount Doom with the One Ring.', 'TheLord.jpg');
INSERT INTO `movie` VALUES (18, 'The Shawshank Redemption', 1994, 93, 'Two imprisoned men bond over a number of years, finding solace and eventual redemption through acts of common decency.', 'theShawshankRedemption.jpg');
INSERT INTO `movie` VALUES (19, 'Titanic', 1997, 79, 'A seventeen-year-old aristocrat falls in love with a kind but poor artist aboard the luxurious, ill-fated R.M.S. Titanic.', 'titanic.jpg');
INSERT INTO `movie` VALUES (20, 'The Wolf of Wall Street', 2013, 3.43333, 'Based on the true story of Jordan Belfort, from his rise to a wealthy stock-broker living the high life to his fall involving crime, corruption and the federal government.', 'wolf.jpg');

-- ----------------------------
-- Table structure for movie_cast
-- ----------------------------
DROP TABLE IF EXISTS `movie_cast`;
CREATE TABLE `movie_cast`  (
  `cast_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  INDEX `cast_id`(`cast_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `movie_cast_ibfk_1` FOREIGN KEY (`cast_id`) REFERENCES `cast` (`cast_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `movie_cast_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie_cast
-- ----------------------------
INSERT INTO `movie_cast` VALUES (1, 1);
INSERT INTO `movie_cast` VALUES (2, 1);
INSERT INTO `movie_cast` VALUES (4, 10);
INSERT INTO `movie_cast` VALUES (6, 2);
INSERT INTO `movie_cast` VALUES (7, 18);
INSERT INTO `movie_cast` VALUES (8, 9);
INSERT INTO `movie_cast` VALUES (9, 3);
INSERT INTO `movie_cast` VALUES (10, 7);
INSERT INTO `movie_cast` VALUES (11, 8);
INSERT INTO `movie_cast` VALUES (12, 11);
INSERT INTO `movie_cast` VALUES (13, 12);
INSERT INTO `movie_cast` VALUES (14, 13);
INSERT INTO `movie_cast` VALUES (15, 14);
INSERT INTO `movie_cast` VALUES (16, 15);
INSERT INTO `movie_cast` VALUES (17, 16);
INSERT INTO `movie_cast` VALUES (18, 17);
INSERT INTO `movie_cast` VALUES (19, 1);
INSERT INTO `movie_cast` VALUES (4, 10);
INSERT INTO `movie_cast` VALUES (4, 19);
INSERT INTO `movie_cast` VALUES (5, 1);
INSERT INTO `movie_cast` VALUES (5, 6);
INSERT INTO `movie_cast` VALUES (3, 1);

-- ----------------------------
-- Table structure for movie_genre
-- ----------------------------
DROP TABLE IF EXISTS `movie_genre`;
CREATE TABLE `movie_genre`  (
  `movie_id` int NULL DEFAULT NULL,
  `genre_id` int NULL DEFAULT NULL,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  INDEX `genre_id`(`genre_id` ASC) USING BTREE,
  CONSTRAINT `movie_genre_ibfk_1` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `movie_genre_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `genre` (`genre_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of movie_genre
-- ----------------------------
INSERT INTO `movie_genre` VALUES (1, 4);
INSERT INTO `movie_genre` VALUES (2, 8);
INSERT INTO `movie_genre` VALUES (3, 4);
INSERT INTO `movie_genre` VALUES (4, 1);
INSERT INTO `movie_genre` VALUES (5, 2);
INSERT INTO `movie_genre` VALUES (6, 4);
INSERT INTO `movie_genre` VALUES (7, 8);
INSERT INTO `movie_genre` VALUES (8, 8);
INSERT INTO `movie_genre` VALUES (9, 8);
INSERT INTO `movie_genre` VALUES (10, 4);
INSERT INTO `movie_genre` VALUES (11, 4);
INSERT INTO `movie_genre` VALUES (12, 8);
INSERT INTO `movie_genre` VALUES (13, 8);
INSERT INTO `movie_genre` VALUES (14, 4);
INSERT INTO `movie_genre` VALUES (15, 4);
INSERT INTO `movie_genre` VALUES (16, 8);
INSERT INTO `movie_genre` VALUES (17, 8);
INSERT INTO `movie_genre` VALUES (18, 8);
INSERT INTO `movie_genre` VALUES (19, 6);
INSERT INTO `movie_genre` VALUES (20, 7);

-- ----------------------------
-- Table structure for rating
-- ----------------------------
DROP TABLE IF EXISTS `rating`;
CREATE TABLE `rating`  (
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  `score` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `rating_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `rating_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of rating
-- ----------------------------
INSERT INTO `rating` VALUES (1, 1, '5');
INSERT INTO `rating` VALUES (1, 5, '4.5');
INSERT INTO `rating` VALUES (1, 2, '5');
INSERT INTO `rating` VALUES (1, 2, '4.5');
INSERT INTO `rating` VALUES (1, 3, '4');
INSERT INTO `rating` VALUES (1, 4, '3.8');
INSERT INTO `rating` VALUES (1, 5, '2.1');
INSERT INTO `rating` VALUES (1, 6, '2.5');
INSERT INTO `rating` VALUES (1, 7, '3.8');
INSERT INTO `rating` VALUES (1, 8, '4.6');
INSERT INTO `rating` VALUES (1, 9, '4.9');
INSERT INTO `rating` VALUES (1, 10, '5.0');
INSERT INTO `rating` VALUES (1, 11, '1.2');
INSERT INTO `rating` VALUES (1, 12, '3.7');
INSERT INTO `rating` VALUES (1, 13, '1.4');
INSERT INTO `rating` VALUES (1, 14, '3.6');
INSERT INTO `rating` VALUES (1, 15, '4.6');
INSERT INTO `rating` VALUES (1, 16, '2.7');
INSERT INTO `rating` VALUES (1, 17, '3.4');
INSERT INTO `rating` VALUES (1, 18, '5.0');
INSERT INTO `rating` VALUES (1, 19, '4.2');
INSERT INTO `rating` VALUES (1, 20, '1.4');
INSERT INTO `rating` VALUES (2, 11, '4.5');
INSERT INTO `rating` VALUES (2, 2, '4.6');
INSERT INTO `rating` VALUES (2, 3, '2.7');
INSERT INTO `rating` VALUES (2, 15, '4.4');
INSERT INTO `rating` VALUES (2, 5, '5.0');
INSERT INTO `rating` VALUES (2, 6, '4.2');
INSERT INTO `rating` VALUES (2, 7, '4.2');
INSERT INTO `rating` VALUES (2, 8, '3.6');
INSERT INTO `rating` VALUES (2, 9, '4.2');
INSERT INTO `rating` VALUES (2, 10, '4.1');
INSERT INTO `rating` VALUES (2, 14, '4.5');
INSERT INTO `rating` VALUES (2, 12, '4.9');
INSERT INTO `rating` VALUES (2, 13, '4.3');
INSERT INTO `rating` VALUES (2, 1, '3.6');
INSERT INTO `rating` VALUES (2, 4, '3.4');
INSERT INTO `rating` VALUES (2, 16, '4.2');
INSERT INTO `rating` VALUES (2, 17, '3.2');
INSERT INTO `rating` VALUES (2, 18, '3.1');
INSERT INTO `rating` VALUES (2, 19, '2.8');
INSERT INTO `rating` VALUES (2, 20, '3.7');
INSERT INTO `rating` VALUES (3, 12, '3.2');
INSERT INTO `rating` VALUES (3, 13, '3.5');
INSERT INTO `rating` VALUES (3, 3, '2.1');
INSERT INTO `rating` VALUES (3, 14, '3.0');
INSERT INTO `rating` VALUES (3, 5, '1.6');
INSERT INTO `rating` VALUES (3, 17, '3.4');
INSERT INTO `rating` VALUES (3, 7, '1.0');
INSERT INTO `rating` VALUES (3, 8, '2.5');
INSERT INTO `rating` VALUES (3, 9, '5.0');
INSERT INTO `rating` VALUES (3, 10, '3.8');
INSERT INTO `rating` VALUES (3, 11, '3.5');
INSERT INTO `rating` VALUES (3, 4, '4.5');
INSERT INTO `rating` VALUES (3, 2, '2.0');
INSERT INTO `rating` VALUES (3, 16, '3.8');
INSERT INTO `rating` VALUES (3, 15, '3.7');
INSERT INTO `rating` VALUES (3, 1, '4.6');
INSERT INTO `rating` VALUES (3, 6, '1.4');
INSERT INTO `rating` VALUES (3, 18, '3.6');
INSERT INTO `rating` VALUES (3, 19, '3.8');
INSERT INTO `rating` VALUES (3, 20, '3.0');
INSERT INTO `rating` VALUES (4, 1, '3.9');
INSERT INTO `rating` VALUES (4, 2, '3.6');
INSERT INTO `rating` VALUES (4, 3, '4.5');
INSERT INTO `rating` VALUES (4, 4, '4.3');
INSERT INTO `rating` VALUES (4, 5, '3.8');
INSERT INTO `rating` VALUES (4, 6, '2.1');
INSERT INTO `rating` VALUES (4, 7, '2.5');
INSERT INTO `rating` VALUES (4, 8, '3.8');
INSERT INTO `rating` VALUES (4, 9, '4.6');
INSERT INTO `rating` VALUES (4, 10, '4.9');
INSERT INTO `rating` VALUES (4, 11, '5.0');
INSERT INTO `rating` VALUES (4, 12, '1.2');
INSERT INTO `rating` VALUES (4, 13, '3.7');
INSERT INTO `rating` VALUES (4, 14, '1.4');
INSERT INTO `rating` VALUES (4, 15, '3.6');
INSERT INTO `rating` VALUES (4, 16, '4.6');
INSERT INTO `rating` VALUES (4, 17, '2.2');
INSERT INTO `rating` VALUES (4, 18, '3.4');
INSERT INTO `rating` VALUES (4, 19, '5.0');
INSERT INTO `rating` VALUES (4, 20, '4.2');
INSERT INTO `rating` VALUES (5, 1, '1.4');
INSERT INTO `rating` VALUES (5, 2, '3.6');
INSERT INTO `rating` VALUES (5, 3, '4.6');
INSERT INTO `rating` VALUES (5, 4, '3.7');
INSERT INTO `rating` VALUES (5, 5, '3.4');
INSERT INTO `rating` VALUES (5, 6, '5.0');
INSERT INTO `rating` VALUES (5, 7, '4.5');
INSERT INTO `rating` VALUES (5, 8, '4.3');
INSERT INTO `rating` VALUES (5, 9, '3.8');
INSERT INTO `rating` VALUES (5, 10, '2.1');
INSERT INTO `rating` VALUES (5, 11, '2.5');
INSERT INTO `rating` VALUES (5, 12, '3.8');
INSERT INTO `rating` VALUES (5, 13, '4.6');
INSERT INTO `rating` VALUES (5, 14, '4.9');
INSERT INTO `rating` VALUES (5, 15, '5.0');
INSERT INTO `rating` VALUES (5, 16, '1.2');
INSERT INTO `rating` VALUES (5, 17, '3.7');
INSERT INTO `rating` VALUES (5, 18, '1.4');
INSERT INTO `rating` VALUES (5, 19, '3.6');
INSERT INTO `rating` VALUES (5, 20, '4.6');
INSERT INTO `rating` VALUES (6, 1, '2.7');
INSERT INTO `rating` VALUES (6, 2, '3.4');
INSERT INTO `rating` VALUES (6, 3, '5.0');
INSERT INTO `rating` VALUES (6, 4, '4.2');
INSERT INTO `rating` VALUES (6, 5, '1.4');
INSERT INTO `rating` VALUES (6, 6, '3.6');
INSERT INTO `rating` VALUES (6, 7, '4.6');
INSERT INTO `rating` VALUES (6, 8, '2.7');
INSERT INTO `rating` VALUES (6, 9, '3.4');
INSERT INTO `rating` VALUES (6, 10, '4.5');
INSERT INTO `rating` VALUES (6, 11, '4.3');
INSERT INTO `rating` VALUES (6, 12, '3.8');
INSERT INTO `rating` VALUES (6, 13, '2.1');
INSERT INTO `rating` VALUES (6, 14, '2.5');
INSERT INTO `rating` VALUES (6, 15, '3.8');
INSERT INTO `rating` VALUES (6, 16, '4.6');
INSERT INTO `rating` VALUES (6, 17, '4.9');
INSERT INTO `rating` VALUES (6, 18, '5.0');
INSERT INTO `rating` VALUES (6, 19, '1.2');
INSERT INTO `rating` VALUES (6, 20, '3.7');

-- ----------------------------
-- Table structure for review
-- ----------------------------
DROP TABLE IF EXISTS `review`;
CREATE TABLE `review`  (
  `review_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  `review_content` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`review_id`) USING BTREE,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `review_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `review_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of review
-- ----------------------------
INSERT INTO `review` VALUES (1, 1, 1, 'This movie is good.');
INSERT INTO `review` VALUES (2, 1, 15, 'I like this movie very much.');
INSERT INTO `review` VALUES (3, 1, 7, 'I do not like this movie.');
INSERT INTO `review` VALUES (4, 2, 4, 'I appreciate the actors\' performance.');
INSERT INTO `review` VALUES (5, 2, 18, 'This movie is worthy in watching. ');
INSERT INTO `review` VALUES (6, 2, 14, 'This movie is a trash.');
INSERT INTO `review` VALUES (7, 3, 11, 'This movie should get an award.');
INSERT INTO `review` VALUES (8, 4, 20, 'This movie is wonderful.');
INSERT INTO `review` VALUES (9, 5, 12, 'I have never watched such an amazing film.');
INSERT INTO `review` VALUES (10, 6, 17, 'This film is awesome.');
INSERT INTO `review` VALUES (11, 6, 6, 'Enjoy watching this movie.');
INSERT INTO `review` VALUES (12, 6, 3, 'Dislike. Worst cast!');
INSERT INTO `review` VALUES (13, 3, 2, 'This movie makes my cry!');
INSERT INTO `review` VALUES (14, 3, 5, 'This movie is really funny!');
INSERT INTO `review` VALUES (15, 4, 8, 'I love the cast so much!');
INSERT INTO `review` VALUES (16, 4, 9, 'Woo-hoo! Big heart to this movie!');
INSERT INTO `review` VALUES (17, 5, 10, 'Love this movie!');
INSERT INTO `review` VALUES (18, 5, 13, 'Trash! Worst!');
INSERT INTO `review` VALUES (19, 3, 16, 'Hate this storyline!');
INSERT INTO `review` VALUES (20, 5, 19, 'Not funny! I donnot wanna watch it again.');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_type` int NULL DEFAULT NULL,
  `gender` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `age` int NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'qwe', '123', 1, 'male', 19, 'erer@gmail.com');
INSERT INTO `user` VALUES (2, 'qwe2', '123456', 2, 'female', 12, '123@qwe.com');
INSERT INTO `user` VALUES (3, 'wangyx', '13579', 3, 'male', 23, 'wangyx@gmail.com');
INSERT INTO `user` VALUES (4, 'blessedroll', '246810', 2, 'male', 29, 'blessedroll@gmail.com');
INSERT INTO `user` VALUES (5, 'dtxpder', '19283', 1, 'female', 23, 'dtxpder@gmail.com');
INSERT INTO `user` VALUES (6, 'bazingaoo', '234667', 1, 'male', 23, 'bazingaoo@gmail.com');

-- ----------------------------
-- Table structure for watch_list
-- ----------------------------
DROP TABLE IF EXISTS `watch_list`;
CREATE TABLE `watch_list`  (
  `user_id` int NULL DEFAULT NULL,
  `movie_id` int NULL DEFAULT NULL,
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `movie_id`(`movie_id` ASC) USING BTREE,
  CONSTRAINT `watch_list_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `watch_list_ibfk_2` FOREIGN KEY (`movie_id`) REFERENCES `movie` (`movie_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of watch_list
-- ----------------------------
INSERT INTO `watch_list` VALUES (1, 2);
INSERT INTO `watch_list` VALUES (1, 15);
INSERT INTO `watch_list` VALUES (1, 7);
INSERT INTO `watch_list` VALUES (2, 4);
INSERT INTO `watch_list` VALUES (2, 18);
INSERT INTO `watch_list` VALUES (2, 14);
INSERT INTO `watch_list` VALUES (3, 11);
INSERT INTO `watch_list` VALUES (4, 20);
INSERT INTO `watch_list` VALUES (5, 12);
INSERT INTO `watch_list` VALUES (6, 17);
INSERT INTO `watch_list` VALUES (6, 6);
INSERT INTO `watch_list` VALUES (6, 3);
INSERT INTO `watch_list` VALUES (3, 2);
INSERT INTO `watch_list` VALUES (3, 5);
INSERT INTO `watch_list` VALUES (4, 8);
INSERT INTO `watch_list` VALUES (4, 9);
INSERT INTO `watch_list` VALUES (5, 10);
INSERT INTO `watch_list` VALUES (5, 13);
INSERT INTO `watch_list` VALUES (3, 16);
INSERT INTO `watch_list` VALUES (5, 19);

SET FOREIGN_KEY_CHECKS = 1;
