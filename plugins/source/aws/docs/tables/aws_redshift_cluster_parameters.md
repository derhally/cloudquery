# Table: aws_redshift_cluster_parameters


The composite primary key for this table is (**cluster_arn**, **parameter_name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|cluster_arn (PK)|String|
|parameter_name (PK)|String|
|allowed_values|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|minimum_engine_version|String|
|parameter_value|String|
|source|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|