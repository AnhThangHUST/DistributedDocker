## CASE 1:
=========
 -If you don't have docker:
 -STEP1: sudo apt-get install docker-ce
 -STEP2: ./compile.sh
 -STEP3: ./stopdocker.sh
 -STEP4: ./startdocker.sh


## CASE 2:

-If you have had docker:
    -If the damemon dies:
        -STEP1:  sudo rm -rf /var/lib/docker
                sudo apt-get purge docker-ce
        -STEP2->STEP5 = STEP1->STEP4(In CASE 1)
    -If the daemon doen't die and Docker Dameon is running:
        STEP1->STEP3 = STEP2->STEP4(In CASE 1)
    -If the daemon doen't die and Docker Dameon isn't running:
        STEP1: ./startdocker.sh
        STEP2->STEP4 = STEP2->STEP4(In CASE 1)
