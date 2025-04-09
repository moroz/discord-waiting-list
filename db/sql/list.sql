-- name: CountWaiters :one
select count(*) from waiting_list;

-- name: InsertWaiter :one
insert into waiting_list (email, name, bio, region)
values (?, ?, ?, ?) returning *;
