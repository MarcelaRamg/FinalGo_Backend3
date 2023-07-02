-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema clinica
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema clinica
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `clinica` DEFAULT CHARACTER SET utf8 ;
USE `clinica` ;

-- -----------------------------------------------------
-- Table `clinica`.`Pacientes`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinica`.`Pacientes` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Nombre` VARCHAR(45) NULL,
  `Apellido` VARCHAR(45) NOT NULL,
  `Dni` DOUBLE NOT NULL,
  `FechaAlta` DATE NOT NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clinica`.`Dentista`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinica`.`Dentista` (
  `ID` INT NOT NULL AUTO_INCREMENT,
  `Apellido` VARCHAR(45) NOT NULL,
  `Nombre` VARCHAR(45) NULL,
  `Matricula` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`ID`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `clinica`.`Turno`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `clinica`.`Turno` (
  `ID` INT NOT NULL,
  `FechaHora` DATETIME NOT NULL,
  `PacienteID` INT NOT NULL,
  `DentistaID` INT NOT NULL,
  `Descripcion` VARCHAR(60) NULL DEFAULT 'Descripcion de ejemplo',
  PRIMARY KEY (`ID`),
  INDEX `idDentista_idx` (`DentistaID` ASC) VISIBLE,
  INDEX `idCliente_idx` (`PacienteID` ASC) VISIBLE,
  CONSTRAINT `idDentista`
    FOREIGN KEY (`DentistaID`)
    REFERENCES `clinica`.`Dentista` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION,
  CONSTRAINT `idCliente`
    FOREIGN KEY (`PacienteID`)
    REFERENCES `clinica`.`Pacientes` (`ID`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
