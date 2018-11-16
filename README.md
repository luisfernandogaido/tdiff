#tdiff

Calcula a diferença entre horários e durações.

##Uso

Calcula a duração entre dois horários:
```
~$ tdiff 15:45 13:31
```

Calcula a duração entre duas datas/horas:
```
~$ tdiff 2018-11-16T15:45 2018-11-15T13:31
```

Calcula a diferença entre um horário e uma duração:

```
~$ tdiff 15:45 43m
```

O horário inicial pode ser omitido se referir à hora atual:
```
~$ tdiff 2h43
``` 
ou:
```
~$ tdiff 13:31
``` 

