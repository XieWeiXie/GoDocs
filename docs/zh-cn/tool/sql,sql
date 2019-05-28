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



1. 注册会员超时
2. 回头客

