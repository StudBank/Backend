
DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'admin') THEN

      RAISE NOTICE 'Role "admin" already exists. Skipping.';
   ELSE
      CREATE ROLE admin LOGIN PASSWORD 'ButterFly777';
   END IF;
END
$do$;


DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'studbank') THEN

      RAISE NOTICE 'Role "studbank" already exists. Skipping.';
   ELSE
      CREATE ROLE studbank LOGIN PASSWORD 'ButterFly777';
   END IF;
END
$do$;

DO
$do$
BEGIN
   IF EXISTS (
      SELECT FROM pg_catalog.pg_roles
      WHERE  rolname = 'ivan') THEN

      RAISE NOTICE 'Role "ivan" already exists. Skipping.';
   ELSE
      CREATE ROLE ivan LOGIN PASSWORD 'ButterFly777';
   END IF;
END
$do$;

