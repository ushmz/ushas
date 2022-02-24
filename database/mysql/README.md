# ushas database


## Requirements

- docker 

## Structure
```
mysql/
  ├── conf.d/
  │  └── # Mysql global configuration files.
  ├── data/
  │  └──  # Database raw data.
  │       # You have to create this directory before build docker image.
  ├── init.d/
  │  └── # Table structure SQL files that exected in initialization.
  └── migrations/
     └── # Migration files handled with "sql-migrate"
```
