会员组

> venus_dev
select group_id from customer_groups where deleted_at is null and group_id !='';


门店版老客

select group_id from customer_groups where company_id in (select id from companies where type='store') and group_type='potential' and group_id !='';

> pandora_dev
select * from groups where name='老客' and is_user=false and device_pack_uuid is null and deleted_at is null;


mall 回头客

> venus_dev
select group_uuid from frequent_customer_groups where company_id in (select id from companies where type='shopping_mall') and deleted_at is  null;







select c.deleted_at, c.person_id, c.customer_group_id, g.group_id from customers as c inner join customer_groups as g on c.deleted_at between '2019-05-23 00:00:00' and '2019-06-13 23:59:59' and c.customer_group_id = g.id;