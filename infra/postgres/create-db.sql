DO
$do$
BEGIN
   IF EXISTS (SELECT FROM pg_database WHERE datname = 'backend') THEN
      RAISE NOTICE 'Database already exists';  -- optional
   ELSE
      PERFORM dblink_exec('dbname=' || current_database()  -- current db
                        , 'CREATE DATABASE backend');
   END IF;
END
$do$;
