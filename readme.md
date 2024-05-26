# Prehnite log

Writes log to a file... 


Logger is the main structure used to manage all the log files.

The logging system is split into 2 parts. 
The Logger contains multiple Binders. The binder is the main way to seperate the logs. ( Similar to a database in SQL Databases)

The binder can be used for individual tests / logs for individual functions. Each binder can be used for the modules. 



The ledger is the controller for the logs directly. ( Similar to a table in SQL Databases).

Logger -< Binder -< Ledger
