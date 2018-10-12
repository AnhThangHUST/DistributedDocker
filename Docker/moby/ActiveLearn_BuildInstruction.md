## Case 1:
If you don't have docker:

  Step 1: sudo apt-get install docker-ce
  
  Step 2: ./compile.sh
  
  Step 3: ./stopdocker.sh


## Case 2:

If you have had docker:

    If the damemon dies:
        
        Step 1:  sudo rm -rf /var/lib/docker
                 
                 sudo apt-get purge docker-ce
        
        Step 2->Step 5 = Step 1->Step 4(In CASE 1)
    
    If the daemon doen't die and Docker Dameon is running:
        
        Step 1->Step 3 = STEP 2->Step 4(In CASE 1)
    
    If the daemon doen't die and Docker Dameon isn't running:
       
       Step 1: ./startdocker.sh
       
       Step 2->Step 4 = Step 2->Step 4(In CASE 1)
