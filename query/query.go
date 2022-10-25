package query

func Get_SSD_Query(querynum string) string {
	switch querynum {
	case "1":
		return `SELECT l_returnflag,
	l_linestatus,
	SUM(l_quantity)                                           AS sum_qty,
	SUM(l_extendedprice)                                      AS
	sum_base_price,
	SUM(l_extendedprice * ( 1 - l_discount ))                 AS
	sum_disc_price,
	SUM(l_extendedprice * ( 1 - l_discount ) * ( 1 + l_tax )) AS sum_charge,
	Avg(l_quantity)                                           AS avg_qty,
	Avg(l_extendedprice)                                      AS avg_price,
	Avg(l_discount)                                           AS avg_disc,
	Count(*)                                                  AS count_order
FROM   lineitem
WHERE  l_shipdate <= DATE ('1998-12-01') - interval '108' day
GROUP  BY l_returnflag,
	   l_linestatus
ORDER  BY l_returnflag,
	   l_linestatus;`
	case "2":
		return `select 
	s_acctbal, 
	s_name, 
	n_name, 
	p_partkey, 
	p_mfgr, 
	s_address, 
	s_phone, 
	s_comment 
from 
	part, 
	supplier, 
	partsupp, 
	nation, 
	region 
where 
	p_partkey = ps_partkey and 
	s_suppkey = ps_suppkey and 
	p_size = 30 and 
	p_type like '%STEEL' and 
	s_nationkey = n_nationkey and
	n_regionkey = r_regionkey and 
	r_name = 'ASIA' and 
	ps_supplycost = (
		select 
			min(ps_supplycost) 
		from 
			partsupp, 
			supplier, 
			nation, 
			region 
		where 
			p_partkey = ps_partkey and 
			s_suppkey = ps_suppkey and 
			s_nationkey = n_nationkey and 
			n_regionkey = r_regionkey and 
			r_name = 'ASIA'
	) 
order by 
	s_acctbal desc, 
	n_name, 
	s_name, 
	p_partkey 
limit 100;
`
	case "3":
		return `SELECT   l_orderkey,
	Sum(l_extendedprice * (1 - l_discount)) AS revenue,
	o_orderdate,
	o_shippriority
FROM     customer,
	orders,
	lineitem
WHERE    c_mktsegment = 'AUTOMOBILE'
AND      c_custkey = o_custkey
AND      l_orderkey = o_orderkey
AND      o_orderdate < date '1995-03-13'
AND      l_shipdate >  date '1995-03-13'
GROUP BY l_orderkey,
	o_orderdate,
	o_shippriority
ORDER BY revenue DESC,
	o_orderdate
LIMIT    10;`
	case "4":
		return `select 
	o_orderpriority, count(*) as order_count 
from 
	orders 
where 
	o_orderdate >= date '1995-01-01' and 
	o_orderdate < date '1995-01-01' + interval '3' month and 
	exists (
		select 
			* 
		from 
			lineitem 
		where 
			l_orderkey = o_orderkey and 
			l_commitdate < l_receiptdate
		) 
group by 
	o_orderpriority 
order by 
	o_orderpriority;
`

	case "5":
		return `select 
	n_name, 
	sum(l_extendedprice * (1 - l_discount)) as revenue 
from 
	customer, 
	orders, 
	lineitem, 
	supplier, 
	nation, 
	region 
where 
	c_custkey = o_custkey and 
	l_orderkey = o_orderkey and 
	l_suppkey = s_suppkey and 
	c_nationkey = s_nationkey and 
	s_nationkey = n_nationkey and 
	n_regionkey = r_regionkey and 
	r_name = 'MIDDLE EAST' and 
	o_orderdate >= date '1994-01-01' and 
	o_orderdate < date '1994-01-01' + interval '1' year
group by 
	n_name 
order by 
	revenue desc;

`
	case "6":
		return `
	select
		sum(l_extendedprice * l_discount) as revenue
	from
		lineitem
	where
		l_shipdate >= date ('1994-01-01')
		and l_shipdate < date ('1994-01-01') + interval '1' year
		and l_discount between 0.06 - 0.01 and 0.06 + 0.01
		and l_quantity < 24;
	`
	case "7":
		return `select 
	supp_nation, 
	cust_nation, 
	l_year, 
	sum(volume) as revenue 
from ( 
	select 
		n1.n_name as supp_nation, 
		n2.n_name as cust_nation, 
		extract(year from l_shipdate) as l_year, 
		l_extendedprice * (1 - l_discount) as volume 
	from 
		supplier, 
		lineitem, 
		orders, 
		customer, 
		nation n1, 
		nation n2 
	where 
		s_suppkey = l_suppkey and 
		o_orderkey = l_orderkey and 
		c_custkey = o_custkey and 
		s_nationkey = n1.n_nationkey and
		c_nationkey = n2.n_nationkey and 
		(
			(n1.n_name = 'JAPAN' and n2.n_name = 'INDIA') or 
			(n1.n_name = 'INDIA' and n2.n_name = 'JAPAN')
		) and 
		l_shipdate between date '1995-01-01' and date '1996-12-31'
	) as shipping 
group by 
	supp_nation, 
	cust_nation, 
	l_year 
order by 
	supp_nation, 
	cust_nation, 
	l_year;
`
	case "8":
		return `select 
	o_year, 
	sum(case when nation = 'INDIA' then volume else 0 end) / sum(volume) as mkt_share 
from (
	select 
		extract(year from o_orderdate) as o_year,	
		l_extendedprice * (1 - l_discount) as volume, 
		n2.n_name as nation 
	from 
		part,
		supplier,
		lineitem,
		orders,
		customer,
		nation n1,
		nation n2,
		region
	where 
		p_partkey = l_partkey and 
		s_suppkey = l_suppkey and 
		l_orderkey = o_orderkey and 
		o_custkey = c_custkey and 
		c_nationkey = n1.n_nationkey and 
		n1.n_regionkey = r_regionkey and 
		r_name = 'ASIA'	and 
		s_nationkey = n2.n_nationkey and 
		o_orderdate between date '1995-01-01' and date '1996-12-31'and 
		p_type = 'SMALL PLATED COPPER'
	) as all_nations 
group by 
	o_year 
order by 
	o_year;`
	case "9":
		return `select 
	nation, 
	o_year, 
	sum(amount) as sum_profit 
from (
	select 
		n_name as nation, 
		extract(year from o_orderdate) as o_year, 
		l_extendedprice * (1 - l_discount) - ps_supplycost * l_quantity as amount 
	from 
		part, 
		supplier, 
		lineitem, 
		partsupp, 
		orders, 
		nation 
	where 
		s_suppkey = l_suppkey and 
		ps_suppkey = l_suppkey and 
		ps_partkey = l_partkey and 
		p_partkey = l_partkey and 
		o_orderkey = l_orderkey and 
		s_nationkey = n_nationkey and 
		p_name like '%dim%'
	) as profit 
group by 
	nation, 
	o_year 
order by 
	nation, 
	o_year desc;`
	case "10":
		return `select c_custkey,
	c_name,
	sum(l_extendedprice * (1 - l_discount)) as revenue,
	c_acctbal,
	n_name,
	c_address,
	c_phone,
	c_comment
from
	customer,
	orders,
	lineitem,
	nation
where
	c_custkey = o_custkey
	and l_orderkey = o_orderkey
	and o_orderdate >= date '1993-08-01'
	and o_orderdate < date '1993-08-01' + interval '3' month
	and l_returnflag = 'R'
	and c_nationkey = n_nationkey
group by
	c_custkey,
	c_name,
	c_acctbal,
	c_phone,
	n_name,
	c_address,
	c_comment
order by
	revenue desc
limit 20;
`
	case "11":
		return `SELECT ps_partkey,
	Sum(ps_supplycost * ps_availqty) AS value
FROM   partsupp,
	supplier,
	nation
WHERE  ps_suppkey = s_suppkey
	AND s_nationkey = n_nationkey
	AND n_name = 'MOZAMBIQUE'
GROUP  BY ps_partkey
HAVING Sum(ps_supplycost * ps_availqty) > (SELECT
	Sum(ps_supplycost * ps_availqty) * 0.0001000000
										FROM   partsupp,
											   supplier,
											   nation
										WHERE  ps_suppkey = s_suppkey
											   AND s_nationkey = n_nationkey
											   AND n_name = 'MOZAMBIQUE')
ORDER  BY value DESC; `
	case "12":
		return `SELECT l_shipmode,
	SUM(CASE
		  WHEN o_orderpriority = '1-urgent'
				OR o_orderpriority = '2-high' THEN 1
		  ELSE 0
		END) AS high_line_count,
	SUM(CASE
		  WHEN o_orderpriority <> '1-urgent'
			   AND o_orderpriority <> '2-high' THEN 1
		  ELSE 0
		END) AS low_line_count
FROM   orders,
	lineitem
WHERE  o_orderkey = l_orderkey
	AND l_shipmode IN ( 'RAIL', 'FOB' )
	AND l_commitdate < l_receiptdate
	AND l_shipdate < l_commitdate
	AND l_receiptdate >= DATE '1997-01-01'
	AND l_receiptdate < DATE '1997-01-01' + interval '1' year
GROUP  BY l_shipmode
ORDER  BY l_shipmode; `
	case "13":
		return `SELECT c_count,
	Count(*) AS custdist
FROM   (SELECT c_custkey,
			Count(o_orderkey) AS c_count
	 FROM   customer
			LEFT OUTER JOIN orders
						 ON c_custkey = o_custkey
							AND o_comment NOT LIKE '%PENDING%DEPOSITS%'
	 GROUP  BY c_custkey) c_orders
GROUP  BY c_count
ORDER  BY custdist DESC,
	   c_count DESC; `
	case "14":
		return `SELECT 100.00 * SUM(CASE
		WHEN p_type LIKE 'PROMO%' THEN l_extendedprice *
									   ( 1 - l_discount )
		ELSE 0
	  END) / SUM(l_extendedprice * ( 1 - l_discount )) AS
promo_revenue
FROM   lineitem,
part
WHERE  l_partkey = p_partkey
AND l_shipdate >= DATE '1996-12-01'
AND l_shipdate < DATE '1996-12-01' + interval '1' month; `
	case "15":
		return `CREATE VIEW revenue0
	(supplier_no, total_revenue)
	AS
	  SELECT l_suppkey,
			 SUM(l_extendedprice * ( 1 - l_discount ))
	  FROM   lineitem
	  WHERE  l_shipdate >= DATE '1997-07-01'
			 AND l_shipdate < DATE '1997-07-01' + interval '3' month
	  GROUP  BY l_suppkey;
	
	SELECT s_suppkey,
		   s_name,
		   s_address,
		   s_phone,
		   total_revenue
	FROM   supplier,
		   revenue0
	WHERE  s_suppkey = supplier_no
		   AND total_revenue = (SELECT Max(total_revenue)
								FROM   revenue0)
	ORDER  BY s_suppkey;
	
	DROP VIEW revenue0; `
	case "16":
		return `SELECT p_brand,
	p_type,
	p_size,
	Count(DISTINCT ps_suppkey) AS supplier_cnt
FROM   partsupp,
	part
WHERE  p_partkey = ps_partkey
	AND p_brand <> 'Brand#34'
	AND p_type NOT LIKE 'LARGE BRUSHED%'
	AND p_size IN ( 48, 19, 12, 4,
					41, 7, 21, 39 )
	AND ps_suppkey NOT IN (SELECT s_suppkey
						   FROM   supplier
						   WHERE  s_comment LIKE '%CUSTOMER%COMPLAINTS%')
GROUP  BY p_brand,
	   p_type,
	   p_size
ORDER  BY supplier_cnt DESC,
	   p_brand,
	   p_type,
	   p_size; `
	case "17":
		return `SELECT SUM(l_extendedprice) / 7.0 AS avg_yearly
	FROM   lineitem,
		   part
	WHERE  p_partkey = l_partkey
		   AND p_brand = 'Brand#44'
		   AND p_container = 'WRAP PKG'
		   AND l_quantity < (SELECT 0.2 * AVG(l_quantity)
							 FROM   lineitem
							 WHERE  l_partkey = p_partkey); `
	case "18":
		return `SELECT c_name,
	c_custkey,
	o_orderkey,
	o_orderdate,
	o_totalprice,
	Sum(l_quantity)
FROM   customer,
	orders,
	lineitem
WHERE  o_orderkey IN (SELECT l_orderkey
				   FROM   lineitem
				   GROUP  BY l_orderkey
				   HAVING Sum(l_quantity) > 314)
	AND c_custkey = o_custkey
	AND o_orderkey = l_orderkey
GROUP  BY c_name,
	   c_custkey,
	   o_orderkey,
	   o_orderdate,
	   o_totalprice
ORDER  BY o_totalprice DESC,
	   o_orderdate
LIMIT  100; `
	case "19":
		return `SELECT Sum(l_extendedprice * ( 1 - l_discount )) AS revenue
	FROM   lineitem,
		   part
	WHERE  ( p_partkey = l_partkey
			 AND p_brand = 'Brand#52'
			 AND p_container IN ( 'SM CASE', 'SM BOX', 'SM PACK', 'SM PKG' )
			 AND l_quantity >= 4
			 AND l_quantity <= 4 + 10
			 AND p_size BETWEEN 1 AND 5
			 AND l_shipmode IN ( 'AIR', 'AIR REG' )
			 AND l_shipinstruct = 'DELIVER IN PERSON' )
			OR ( p_partkey = l_partkey
				 AND p_brand = 'Brand#11'
				 AND p_container IN ( 'MED bag', 'MED BOX', 'MED PKG', 'MED PACK' )
				 AND l_quantity >= 18
				 AND l_quantity <= 18 + 10
				 AND p_size BETWEEN 1 AND 10
				 AND l_shipmode IN ( 'AIR', 'AIR REG' )
				 AND l_shipinstruct = 'DELIVER IN PERSON' )
			OR ( p_partkey = l_partkey
				 AND p_brand = 'Brand#51'
				 AND p_container IN ( 'LG case', 'LG BOX', 'LG PACK', 'LG PKG' )
				 AND l_quantity >= 29
				 AND l_quantity <= 29 + 10
				 AND p_size BETWEEN 1 AND 15
				 AND l_shipmode IN ( 'AIR', 'AIR REG' )
				 AND l_shipinstruct = 'DELIVER IN PERSON' ); `
	case "20":
		return `SELECT s_name,
	s_address
FROM   supplier,
	nation
WHERE  s_suppkey IN (SELECT ps_suppkey
				  FROM   partsupp
				  WHERE  ps_partkey IN (SELECT p_partkey
										FROM   part
										WHERE  p_name LIKE 'green%')
						 AND ps_availqty > (SELECT 0.5 * SUM(l_quantity)
											FROM   lineitem
											WHERE  l_partkey = ps_partkey
												   AND l_suppkey = ps_suppkey
												   AND l_shipdate >= DATE
													   '1993-01-01'
												   AND l_shipdate < DATE
													   '1993-01-01' +
													   interval
																	'1' year
										   ))
	AND s_nationkey = n_nationkey
	AND n_name = 'ALGERIA'
ORDER  BY s_name; `
	case "21":
		return `SELECT s_name,
	Count(*) AS numwait
FROM   supplier,
	lineitem l1,
	orders,
	nation
WHERE  s_suppkey = l1.l_suppkey
	AND o_orderkey = l1.l_orderkey
	AND o_orderstatus = 'F'
	AND l1.l_receiptdate > l1.l_commitdate
	AND EXISTS (SELECT *
				FROM   lineitem l2
				WHERE  l2.l_orderkey = l1.l_orderkey
					   AND l2.l_suppkey <> l1.l_suppkey)
	AND NOT EXISTS (SELECT *
					FROM   lineitem l3
					WHERE  l3.l_orderkey = l1.l_orderkey
						   AND l3.l_suppkey <> l1.l_suppkey
						   AND l3.l_receiptdate > l3.l_commitdate)
	AND s_nationkey = n_nationkey
	AND n_name = 'EGYPT'
GROUP  BY s_name
ORDER  BY numwait DESC,
	   s_name
LIMIT  100; `
	case "22":
		return `SELECT cntrycode,
	Count(*)       AS numcust,
	Sum(c_acctbal) AS totacctbal
FROM   (SELECT Substring(c_phone FROM 1 FOR 2) AS cntrycode,
			c_acctbal
	 FROM   customer
	 WHERE  Substring(c_phone FROM 1 FOR 2) IN ( '20', '40', '22', '30',
												 '39', '42', '21' )
			AND c_acctbal > (SELECT Avg(c_acctbal)
							 FROM   customer
							 WHERE  c_acctbal > 0.00
									AND Substring(c_phone FROM 1 FOR 2) IN (
										'20', '40', '22', '30',
										'39', '42', '21' ))
			AND NOT EXISTS (SELECT *
							FROM   orders
							WHERE  o_custkey = c_custkey)) AS custsale
GROUP  BY cntrycode
ORDER  BY cntrycode; `
	}
	return ""
}
