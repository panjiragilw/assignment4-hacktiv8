PGDMP     0    #                  z            asg-4 #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1) #   12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    16503    asg-4    DATABASE     y   CREATE DATABASE "asg-4" WITH TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'en_US.UTF-8' LC_CTYPE = 'en_US.UTF-8';
    DROP DATABASE "asg-4";
                postgres    false            ?            1259    16506    todos    TABLE     ?   CREATE TABLE public.todos (
    id integer NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    completed boolean NOT NULL
);
    DROP TABLE public.todos;
       public         heap    postgres    false            ?            1259    16504    todos_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.todos_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.todos_id_seq;
       public          postgres    false    203            ?           0    0    todos_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.todos_id_seq OWNED BY public.todos.id;
          public          postgres    false    202                       2604    16509    todos id    DEFAULT     d   ALTER TABLE ONLY public.todos ALTER COLUMN id SET DEFAULT nextval('public.todos_id_seq'::regclass);
 7   ALTER TABLE public.todos ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    203    202    203            ?          0    16506    todos 
   TABLE DATA           B   COPY public.todos (id, title, description, completed) FROM stdin;
    public          postgres    false    203   ?
       ?           0    0    todos_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.todos_id_seq', 6, true);
          public          postgres    false    202                       2606    16514    todos todos_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.todos
    ADD CONSTRAINT todos_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.todos DROP CONSTRAINT todos_pkey;
       public            postgres    false    203            ?   X   x?3?t?????K??
????
i??)Ŝi\????%?
!E??
Ήy??9??y
?
%?
%`?d?@b^?BqANf	B???+F??? w? X     