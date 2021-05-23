# Removes the known host from ~/.ssh/known_hosts
from pathlib import Path

known_hosts_path = str(Path.home()) + "/.ssh/known_hosts"

with open(known_hosts_path) as hosts_file:
    lines = hosts_file.readlines()

for i in range(len(lines)):
    if lines[i].startswith("[localhost]:2222"):
        del lines[i]
        print("Removed host")


with open(known_hosts_path, "w") as hosts_file:
    hosts_file.writelines(lines)
