CREATE TABLE nodes (
    id serial PRIMARY KEY,
    host character varying(255),
    free character varying(255),
    alive boolean default true,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE INDEX nodes_alive_index ON nodes(alive);

CREATE TABLE objects (
    id bigserial PRIMARY KEY,
    node_id integer,
    url text,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

CREATE INDEX objects_node_id_index ON objects(node_id);
CREATE INDEX objects_url_index ON objects(url);
