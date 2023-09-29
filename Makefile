task-01:
	go run 01_action_embeds_human/main.go

task-02:
	go run 02_concurrent_squares_stdout/main.go

task-03:
	go run 03_concurrent_squares_sum/main.go

# программа подразумевает аргумент коммандной строки при запуске
task-04:
	go run 04_workers_to_stdout/main.go 10

# программа подразумевает аргумент коммандной строки при запуске
task-05:
	go run 05_channel_transmitting_n_second/main.go 5

task-06:
	go run 06_stop_goroutines/main.go

task-07:
	go run 07_concurrent_map_write/main.go
task-07-test:
	go test ./07_concurrent_map_write/... -race

task-08:
	go run 08_set_ith_bit/main.go

task-09:
	go run 09_numbers_convey/main.go

task-10:
	go run 10_temperature_group/main.go

task-11:
	go run 11_sets_intersect/main.go

task-12:
	go run 12_set_for_strings/main.go

task-13:
	go run 13_change_position/main.go

task-14:
	go run 14_identify_type/main.go

task-15:
	go run 15_huge_string/main.go

task-16:
	go run 16_quick_sort/main.go

task-17:
	go run 17_binary_search/main.go

task-18:
	go run 18_concurrent_counter/main.go
task-18-test:
	go test ./18_concurrent_counter/... -race

task-19:
	go run 19_reverse_string/main.go

task-20:
	go run 20_reverse_words_in_string/main.go

task-21:
	go run 21_adapter/main.go

task-22:
	go run 22_big_numbers/main.go

task-23:
	go run 23_delete_from_slice/main.go

task-24:
	go run 24_distance_to_point/main.go

task-25:
	go run 25_sleep/main.go

task-26:
	go run 26_unique_symbols/main.go



# P.S. используешь VScode и тебе надоело набирать с клавиатуры make что-то-там?
# При этом хочется чтобы таргеты были ясные, без сокращений в духе rdsd?
#
# Тогда предлагаю тебе установить расширение, которое добавит codelens для запуска по клику!
# https://marketplace.visualstudio.com/items?itemName=hablof.makefile-buttons