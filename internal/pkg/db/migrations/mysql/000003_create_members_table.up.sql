CREATE TABLE IF NOT EXISTS Members(
    ID INT NOT NULL UNIQUE AUTO_INCREMENT,
    Name VARCHAR (127) NOT NULL ,
    Age INT,
    PRIMARY KEY (ID)
)