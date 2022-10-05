DROP TABLE IF EXISTS systems;

DROP TABLE IF EXISTS groups;

DROP TABLE IF EXISTS repos;

DROP TABLE IF EXISTS reviews;

DROP TABLE IF EXISTS pull_requests;

DROP TABLE IF EXISTS cd_pipelines;

DROP TABLE IF EXISTS deployments;

CREATE TABLE systems (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name VARCHAR (20) UNIQUE,
  type VARCHAR (20) NOT NULL -- TODO: Consider if this should be its own table
);

CREATE TABLE groups (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  system_id INT,
  name VARCHAR (255) NOT NULL,
  description VARCHAR (255),
  key VARCHAR (255),
  UNIQUE (system_id, name),
  CONSTRAINT fk_system FOREIGN KEY(system_id) REFERENCES systems(id)
);

CREATE TABLE repos (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  group_id INT,
  scm_id INT,
  name VARCHAR (255) NOT NULL,
  slug VARCHAR (255) NOT NULL,
  UNIQUE (scm_id, group_id),
  CONSTRAINT fk_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE pull_requests (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  repo_id INT,
  number INT NOT NULL,
  source_ref VARCHAR (255) NOT NULL,
  target_ref VARCHAR (255) NOT NULL,
  closed BOOLEAN NOT NULL,
  merged BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  closed_at BOOLEAN NOT NULL,
  UNIQUE (number, repo_id),
  CONSTRAINT fk_repo FOREIGN KEY(repo_id) REFERENCES repos(id)
);

CREATE TABLE reviews (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  pr_id INT,
  created_at TIMESTAMP NOT NULL,
  CONSTRAINT fk_pr FOREIGN KEY(pr_id) REFERENCES pull_requests(id)
);

CREATE TABLE cd_pipelines (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  group_id INT,
  name VARCHAR (255) NOT NULL,
  UNIQUE (name, group_id),
  CONSTRAINT fk_group FOREIGN KEY(group_id) REFERENCES groups(id)
);

CREATE TABLE deployment (
  id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  cd_pipeline_id INT,
  started_at TIMESTAMP NOT NULL,
  ended_at TIMESTAMP NOT NULL,
  status VARCHAR (32),
  UNIQUE (cd_pipeline_id, started_at),
  CONSTRAINT fk_cd_pipeline FOREIGN KEY(cd_pipeline_id) REFERENCES cd_pipelines(id)
);

-- insert systems
INSERT INTO
  systems (name, type)
VALUES
  ('github', 'git');

INSERT INTO
  systems (name, type)
VALUES
  ('bitbucket', 'git');

INSERT INTO
  systems (name, type)
VALUES
  ('octopus', 'cd');