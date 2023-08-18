-- Remove null descriptions
UPDATE `languages` SET `desc` = NULL WHERE `languages`.`desc` = '';
