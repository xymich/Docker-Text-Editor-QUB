import requests
import time
import logging

logging.basicConfig(filename='editor_monitor.log', 
                    level=logging.INFO, 
                    format='%(asctime)s - %(levelname)s - %(message)s')

services = {
    "scramble": "http://localhost:8010/scramble?text=test",
    "reverse": "http://localhost:8020/reverse?text=test",
    "duplicate": "http://localhost:8030/duplicate?text=test",
    "commacount": "http://localhost:8040/commacount?text=test"
}

def check_service(name, url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            logging.info(f"{name} service is up. Response: {response.json()}")
        else:
            print("Severe issue with " + name + ", Please check logs.\n")
            logging.error(f"{name} service is down. Status code: {response.status_code}")
    except requests.RequestException as e:
        print("Severe issue with " + name + ", Please check logs.\n")
        logging.error(f"{name} service is down. Error: {e}")

def main():
    while True:
        for name, url in services.items():
            check_service(name, url)
        # Wait 20 seconds
        print("sleeping... (20 seconds)\n")
        time.sleep(20)

if __name__ == '__main__':
    main()
