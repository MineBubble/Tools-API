create database tools_api;

create table tools(
	id serial,
	title varchar(64),
	link varchar(255),
	description varchar(512),
	tags text[]
);

insert into tools(title, link, description, tags) 
values ('Notion',
		'https://notion.so',
		'All in one tool to organize teams and ideas. Write, plan, collaborate, and get organized.', 
		ARRAY ['organization',
			   'planning',
			   'collaboration',
			   'writing',
			   'calendar']),
		('json-server',
		 'https://github.com/typicode/json-server',
		 'Fake REST API based on a json schema. Useful for mocking and creating APIs for front-end devs to consume in coding challenges.',
		 ARRAY ['api',
			   'json',
			   'schema',
			   'node',
			   'github',
			   'rest']),
		('fastify',
		'https://www.fastify.io/',
		'Extremely fast and simple, low-overhead web framework for NodeJS. Supports HTTP2.',
		ARRAY ['web',
			  'fremework',
			  'node',
			  'http2',
			  'https',
			  'localhost']);
		
select * from tools;