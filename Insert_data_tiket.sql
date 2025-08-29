insert into tiket_bioskop_2.users (email,password_hash,full_name,role)
values ('admin@example.com',crypt('admin123', gen_salt('bf')),'Administrator','Admin'),('user1@example.com',crypt('password123', gen_salt('bf')),'Ana Test','User');

insert into tiket_bioskop_2.movies (title, duration_minutes)
values
('Avengers: Endgame', 181),
('Inception', 148),
('The Batman', 176),
('Incredibles 2', 145),
('Bleach Thousand Years', 99);

insert into tiket_bioskop_2.theaters (name, city)
values
('CGV Grand Indonesia', 'Jakarta'),
('XXI Tunjungan Plaza', 'Surabaya'),
('Cinepolis Paragon Mall', 'Semarang');

insert into tiket_bioskop_2.showtimes (movie_id, theater_id, start_at, base_price)
select m.id, t.id, '2025-08-30 19:00:00', 50000
from tiket_bioskop_2.movies m, tiket_bioskop_2.theaters t
where m.title = 'Avengers: Endgame' and t.name = 'CGV Grand Indonesia';

insert into tiket_bioskop_2.showtimes (movie_id, theater_id, start_at, base_price)
select m.id, t.id, '2025-08-30 21:30:00', 45000
from tiket_bioskop_2.movies m, tiket_bioskop_2.theaters t
where m.title = 'Inception' and  t.name = 'XXI Tunjungan Plaza';

insert into tiket_bioskop_2.showtimes (movie_id, theater_id, start_at, base_price)
select m.id, t.id, '2025-08-31 20:00:00', 48000
from tiket_bioskop_2.movies m, tiket_bioskop_2.theaters t
where m.title = 'The Batman' and t.name = 'Cinepolis Paragon Mall';

insert into tiket_bioskop_2.seats (showtime_id, seat_code)
select sh.id, s.seat_code
from tiket_bioskop_2.showtimes sh
cross join (
  values ('A1'), ('A2'), ('A3'), ('A4'), ('A5')
) as s(seat_code);