def main():
    invalid_sum = 0
    id_ranges = open('day_2_input').readline().strip().split(',')
    for id_range in id_ranges:
        id_low, id_high = id_range.split('-')
        for id in range(int(id_low), int(id_high)+1):
            id_str = str(id)
            for segment_length in range(len(id_str)//2, 0, -1):
                segments = [id_str[i:i+segment_length] for i in range(0, len(id_str), segment_length)]
                if len(set(segments)) == 1: 
                    # print(id)
                    invalid_sum += id
                    break
                
    print(invalid_sum)


if __name__ == '__main__':
    main()