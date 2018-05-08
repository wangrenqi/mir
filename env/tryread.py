# coding:utf8

def get_map_ctn():
    ctn = ''
    with open('/Users/Mccree/gopath/src/mir/tmp/readmap/0.map', 'rb') as f:
        ctn = f.read()
    return ctn


def read_int16(bytes):
    big = bytes[1]
    little = bytes[0]
    return little|big<<8


def read_int32(bytes):
    return bytes[3] << 24 | bytes[2] << 16 | bytes[1] << 8 | bytes[0]


def main():
    ctn = get_map_ctn()
    offset = 21
    # print('%s' % ctn[offset:100])
    w = read_int16(ctn[offset: offset+2])
    xor = read_int16(ctn[offset+2: offset+4])
    width = w ^ xor
    print(width)
    h = read_int16(ctn[offset+4: offset+6])
    height = h ^ xor
    print(height)

    offset = 54
    s = []
    for _ in range(width):
        for _ in range(width):
            if ((read_int32(ctn[offset: offset+4]) ^ 0xAA38AA38) & 0x20000000) != 0:
                s.append('.')
            elif ((read_int16(ctn[offset+6: offset+8]) ^ xor) & 0x8000) != 0:
                s.append(' ')
            else:
                s.append(' ')
            offset = offset + 15
        s.append('\n')

    with open('res', 'w') as f:
        f.write(''.join(s))


if __name__ == '__main__':
    main()
