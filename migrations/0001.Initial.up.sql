CREATE TABLE public.applications (
    id bigint DEFAULT next_id() NOT NULL,
    description text NOT NULL,
    is_active boolean NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    CONSTRAINT application_pkey PRIMARY KEY (id)
);

CREATE TABLE public.metadata (
     id bigint DEFAULT next_id() NOT NULL,
     application_id bigint NOT NULL,
     name text NOT NULL,
     owner text NOT NULL,
     configuration_manager text NOT NULL,
     updated_at timestamp with time zone DEFAULT now() NOT NULL,
     created_at timestamp with time zone DEFAULT now() NOT NULL,
     CONSTRAINT metadata_pkey PRIMARY KEY (id),
     CONSTRAINT metadata_application_fkey FOREIGN KEY (application_id)
         REFERENCES public.applications (id) MATCH SIMPLE
         ON UPDATE NO ACTION
         ON DELETE NO ACTION

);

CREATE TABLE public.metadata_history (
     id bigint DEFAULT next_id() NOT NULL,
     application_id bigint NOT NULL,
     metadata_id bigint NOT NULL,
     name text NOT NULL,
     owner text NOT NULL,
     version int NOT NULL,
     configuration_manager text NOT NULL,
     updated_at timestamp with time zone DEFAULT now() NOT NULL,
     created_at timestamp with time zone DEFAULT now() NOT NULL,
     CONSTRAINT metadata_history_pkey PRIMARY KEY (id),
     CONSTRAINT metadata_history_application_fkey FOREIGN KEY (application_id)
         REFERENCES public.applications (id) MATCH SIMPLE
         ON UPDATE NO ACTION
         ON DELETE NO ACTION,
     CONSTRAINT metadata_history_metadata_fkey FOREIGN KEY (metadata_id)
         REFERENCES public.metadata (id) MATCH SIMPLE
         ON UPDATE NO ACTION
         ON DELETE NO ACTION

);