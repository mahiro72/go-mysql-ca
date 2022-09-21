CREATE TABLE `tasks` (
  `id`     int(11) COLLATE utf8mb4_bin NOT NULL COMMENT 'task id',
  `title`  varchar(255) COLLATE utf8mb4_bin NOT NULL COMMENT 'taskのtitle',
  `done`   boolean COLLATE utf8mb4_bin NOT NULL COMMENT 'taskが完了したかどうか',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='taskを管理するtable';

