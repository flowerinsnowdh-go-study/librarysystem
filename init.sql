/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
DROP TABLE IF EXISTS `student`;
DROP TABLE IF EXISTS `book`;

CREATE TABLE `student` (
    `id`   INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` VARCHAR(16) NOT NULL
);
CREATE TABLE `book` (
    `id`          INTEGER PRIMARY KEY AUTOINCREMENT,
    `name`        VARCHAR(16) NOT NULL,
    `borrowed_by` INT,
    FOREIGN KEY (`borrowed_by`) REFERENCES `student`(`id`)
);

-- 添加示例数据
INSERT INTO `student` (`name`) VALUES ('张三');
INSERT INTO `student` (`name`) VALUES ('李四');
INSERT INTO `student` (`name`) VALUES ('王五');
INSERT INTO `student` (`name`) VALUES ('赵六');

INSERT INTO `book` (`name`, `borrowed_by`) VALUES ('挪威的森林', 2);
INSERT INTO `book` (`name`) VALUES ('飞鸟集');
INSERT INTO `book` (`name`) VALUES ('小王子');
INSERT INTO `book` (`name`) VALUES ('朝花夕拾');
INSERT INTO `book` (`name`) VALUES ('月亮与六便士');
