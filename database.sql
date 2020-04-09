create table stats(
  node_id varchar(30) primary key,
  community_chain varchar(30),
	total_memory int,
  used_memory  int,
  used_memory_percentage decimal(10,2),
  processors text,
  most_used_processor_percentage decimal(10,2),
	net_interfaces text,
	most_used_interface_income_bytes bigint,
	most_used_interface_outcome_bytes bigint
);