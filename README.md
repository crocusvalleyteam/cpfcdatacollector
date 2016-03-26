

# project moved to thecroydonproject (https://github.com/thecroydonproject) acccount


# cpfcdatacollector
a html parser to collect match data from public web site

 
 
#introduction
A GO module to collect crystal palace results from various public sites. The programme will be used to build the database

for the croydon project.   

#logic  

 1) using ``http.GET`` load parse the target page     
 2) tokenise the ``resp.Body`` using '"golang.org/x/net/html"``     
 3) walk through the tags to identify table cells ``td``  
 4) processes the content sections of table cells  
 
 
 
#to do

	1) More Testing
	2) Dockerise
	4) Documentation 
	5) Automation and bulk upload of project data


 ```EEEEEEEEEEEEEEEgles```  not a typo but
