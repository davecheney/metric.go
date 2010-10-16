DIRS=agent

all: agent

agent: force_look
	cd agent; $(MAKE) $(MFLAGS)

clean:
	-for d in $(DIRS); do (cd $$d; $(MAKE) clean ); done
	
force_look: 
	-true