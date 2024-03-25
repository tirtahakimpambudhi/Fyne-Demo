FILE_NAME_START_FOLDER = helloworld label button

start_run: $(addprefix build_,$(FILE_NAME_START_FOLDER))
start_run_%:
	go run ./start/$*.go