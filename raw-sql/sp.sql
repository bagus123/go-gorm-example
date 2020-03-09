CREATE OR REPLACE FUNCTION public.getpeople()
 RETURNS SETOF peoples
 LANGUAGE plpgsql
AS $function$
	BEGIN
        return QUERY select * from peoples;
	END;
$function$;


CREATE OR REPLACE FUNCTION public.math_add(x integer, y integer, OUT z integer)
 RETURNS integer
 LANGUAGE plpgsql
AS $function$
	begin
		z := x+y;
	END;
$function$;


CREATE OR REPLACE FUNCTION public.math_add2(x integer, y integer)
 RETURNS TABLE(w integer, z integer)
 LANGUAGE plpgsql
AS $function$
	begin
		w := x + y;
	    z:=w+y;
		return QUERY select w, z;

	END;
$function$;


