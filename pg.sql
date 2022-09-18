DROP TABLE IF EXISTS systems;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS repositories;
DROP TABLE IF EXISTS pull_requests;
DROP TABLE IF EXISTS reviews;
DROP TABLE IF EXISTS deployments;

CREATE TABLE systems (
    name VARCHAR (20),
    type VARCHAR (20) NOT NULL,
    PRIMARY KEY (name)
);

CREATE TABLE groups (
    system VARCHAR (20),
    name VARCHAR (255) NOT NULL,
    PRIMARY KEY (system, name),
    CONSTRAINT fk_system
      FOREIGN KEY(system) 
	  REFERENCES system(customer_id)
);

INSERT INTO systems (name, type) VALUES ('github', 'git');
INSERT INTO systems (name, type) VALUES ('bitbucket', 'git');


