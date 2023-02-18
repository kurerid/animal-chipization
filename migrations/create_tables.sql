CREATE TABLE "Animal"
(
    "Id"               serial       NOT NULL,
    "Weight"           FLOAT        NOT NULL,
    "Length"           FLOAT        NOT NULL,
    "Height"           FLOAT        NOT NULL,
    "Gender"           varchar(255) NOT NULL,
    "LifeStatus"       varchar(255) NOT NULL,
    "Death"            timestamp,
    "ChippingDate"     TIMESTAMP    NOT NULL,
    "Chipper"          integer      NOT NULL,
    "ChippingLocation" integer      NOT NULL,
    CONSTRAINT "Animal_pk" PRIMARY KEY ("Id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "Account"
(
    "Id"        serial       NOT NULL,
    "FirstName" varchar(255) NOT NULL,
    "LastName"  varchar(255) NOT NULL,
    "Email"     varchar(255) NOT NULL,
    "Password"  varchar(255) NOT NULL,
    CONSTRAINT "Account_pk" PRIMARY KEY ("Id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "Type"
(
    "Id"   serial       NOT NULL,
    "Type" varchar(255) NOT NULL,
    CONSTRAINT "Type_pk" PRIMARY KEY ("Id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "Location"
(
    "Id"        serial NOT NULL,
    "Latitude"  FLOAT  NOT NULL,
    "Longitude" FLOAT  NOT NULL,
    CONSTRAINT "Location_pk" PRIMARY KEY ("Id")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "VisitedLocation"
(
    "Animal"      integer   NOT NULL,
    "Location"    integer   NOT NULL,
    "TimeOfVisit" TIMESTAMP NOT NULL,
    CONSTRAINT "VisitedLocation_pk" PRIMARY KEY ("Animal")
) WITH (
      OIDS= FALSE
    );



CREATE TABLE "AnimalType"
(
    "Type"   integer NOT NULL,
    "Animal" integer NOT NULL,
    CONSTRAINT "AnimalType_pk" PRIMARY KEY ("Type", "Animal")
) WITH (
      OIDS= FALSE
    );



ALTER TABLE "Animal"
    ADD CONSTRAINT "Animal_fk0" FOREIGN KEY ("Chipper") REFERENCES "Account" ("Id");
ALTER TABLE "Animal"
    ADD CONSTRAINT "Animal_fk1" FOREIGN KEY ("ChippingLocation") REFERENCES "Location" ("Id");



ALTER TABLE "VisitedLocation"
    ADD CONSTRAINT "VisitedLocation_fk0" FOREIGN KEY ("Animal") REFERENCES "Animal" ("Id");
ALTER TABLE "VisitedLocation"
    ADD CONSTRAINT "VisitedLocation_fk1" FOREIGN KEY ("Location") REFERENCES "Location" ("Id");

ALTER TABLE "AnimalType"
    ADD CONSTRAINT "AnimalType_fk0" FOREIGN KEY ("Type") REFERENCES "Type" ("Id");
ALTER TABLE "AnimalType"
    ADD CONSTRAINT "AnimalType_fk1" FOREIGN KEY ("Animal") REFERENCES "Animal" ("Id");