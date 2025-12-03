def main():
    valid_sum = 0
    id_ranges = open('day_2_input').readline().strip().split(',')
    for id_range in id_ranges:
        id_low, id_high = id_range.split('-')
        if len(id_low) == len (id_high) and len(id_low) % 2 == 1: continue
        for id in range(int(id_low), int(id_high)+1):
            id_str = str(id)
            if len(id_str) % 2 == 1: continue
            if id_str[:len(id_str)//2] == id_str[len(id_str)//2:]: valid_sum += id
    print(valid_sum)


if __name__ == '__main__':
    main()